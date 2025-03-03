<script lang="ts" setup>
const props = withDefaults(
  defineProps<{
    text: string
    href: string
    underline?: boolean
  }>(),
  {
    underline: false
  }
)

const router = useRouter()

function navigateToRoute(event: MouseEvent) {
  if (event.ctrlKey || event.metaKey) {
    window.open(props.href, "_blank")
  } else {
    router.push(props.href)
  }
}

const classList = computed(() => {
  return props.underline ? ["link-style", "underline"] : ["link-style"]
})
</script>

<template>
  <a :class="classList" :href="props.href" @click.prevent="navigateToRoute">{{ props.text }}</a>
</template>

<style lang="scss" scoped>
.link-style {
  display: inline-block;
  max-width: 100%;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  vertical-align: middle;
  line-height: 1.5;
  text-decoration: none;
  color: var(--el-color-primary);

  &:hover {
    opacity: 0.7;
    cursor: pointer;
  }
}

.underline:hover {
  text-decoration: underline;
  text-decoration-skip-ink: none;
  text-underline-offset: 4px;
}
</style>

>
