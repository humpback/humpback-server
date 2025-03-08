import { ServiceVolumeType } from "@/models"
import { NewServiceMetaDockerEmptyInfo, RegistryInfo, ServiceMetaDockerInfo } from "@/types"
import { filter, find, map, omit, omitBy, startsWith } from "lodash-es"
import { GenerateUUID } from "@/utils"

export interface ServiceApplicationInfo extends ServiceMetaDockerInfo {
  imageDomain: string
  imageName: string
  validEnv: Array<{ id: string; name: string; value: string }>
  validLabel: Array<{ id: string; name: string; value: string }>
  validVolumes: Array<{
    id: string
    type: ServiceVolumeType.VolumeTypeBind | ServiceVolumeType.VolumeTypeVolume
    target: string
    source: string
    "readonly": boolean
  }>
  validLogConfig: { type: string; options: Array<{ id: string; name: string; value: string }> }
  validPorts: Array<{
    id: string
    hostPort?: number
    containerPort?: number
    protocol: string
  }>
}

export function NewApplicationInfo(registries: RegistryInfo[], info?: ServiceMetaDockerInfo): ServiceApplicationInfo {
  const validData = omitBy(info, (value, key) => value === undefined || value === null)
  const metaInfo = Object.assign({}, NewServiceMetaDockerEmptyInfo(), validData)
  const imageInfo = parseImageInfo(registries, metaInfo.image)
  return {
    ...metaInfo,
    imageDomain: imageInfo.domain,
    imageName: imageInfo.imageName,
    validEnv: filter(
      map(metaInfo.env, x => {
        const s = x.split("=")
        return {
          id: GenerateUUID(),
          name: s.length > 0 ? s[0] : "",
          value: s.length > 1 ? s[1] : ""
        }
      }),
      d => !!d.name && !!d.value
    ),
    validLabel: map(Object.keys(metaInfo.labels), x => ({ id: GenerateUUID(), name: x, value: metaInfo.labels[x] })),
    validVolumes: map(metaInfo.volumes, x => ({
      id: GenerateUUID(),
      type: x.type,
      target: x.target,
      source: x.source,
      readonly: x.readonly
    })),
    validLogConfig: {
      type: metaInfo.logConfig!.type,
      options: map(Object.keys(metaInfo.logConfig!.config), x => ({
        id: GenerateUUID(),
        name: x,
        value: metaInfo.logConfig!.config[x]
      }))
    },
    validPorts: map(metaInfo.network!.ports, x => ({
      id: GenerateUUID(),
      containerPort: x.containerPort || undefined,
      protocol: x.protocol,
      hostPort: x.hostPort || undefined
    }))
  }
}

function parseImageInfo(registries: RegistryInfo[], image: string) {
  const defaultImage = find(registries, x => x.isDefault)?.url || ""
  const matchRegistryUrls = map(
    filter(registries, x => startsWith(image, x.url)),
    r => r.url
  )
  const matchImageDomain = matchRegistryUrls.reduce((a, b) => (a.length > b.length ? a : b), "")
  if (image) {
    return {
      domain: matchImageDomain || defaultImage,
      imageName: image.replace(`${matchImageDomain}/`, "")
    }
  }
  return {
    domain: defaultImage || "docker.io",
    imageName: ""
  }
}

export function ParseMetaInfo(info: ServiceApplicationInfo): ServiceMetaDockerInfo {
  info.imageName = info.imageName.replace(/\/+/g, "/").replace(/^\/|\/$/g, "")
  return {
    image: `${info.imageDomain}/${info.imageName}`,
    alwaysPull: info.alwaysPull,
    command: info.command,
    env: map(info.validEnv, x => `${x.name}=${x.value}`) || [],
    labels: parseArrayToMap(info.validLabel),
    privileged: info.privileged,
    capabilities: info.capabilities,
    logConfig: {
      type: info.validLogConfig.type,
      config: parseArrayToMap(info.validLogConfig.options)
    },
    resources: info.resources,
    volumes: map(info.validVolumes, x => omit(x, ["id"])) || [],
    network: {
      mode: info.network!.mode,
      hostname: info.network!.hostname,
      networkName: info.network!.networkName,
      useMachineHostname: info.network!.useMachineHostname,
      ports:
        map(info.validPorts, x => ({
          hostPort: x.hostPort || 0,
          containerPort: x.containerPort || 0,
          protocol: x.protocol
        })) || []
    },
    restartPolicy: info.restartPolicy
  }
}

function parseArrayToMap(arr: any[]) {
  const result: any = {}
  map(arr, x => {
    result[x.name] = x.value
  })
  return result
}
