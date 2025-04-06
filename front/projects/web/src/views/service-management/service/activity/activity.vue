<script lang="ts" setup>
import { QueryInfo } from "@/types"
import { TableHeight } from "@/utils"

const { t } = useI18n()

const tableHeight = computed(() => TableHeight(362))

const oldData = ref({
  "dbId": 40,
  "orgId": "ibwirY7jtPyFY4JmEZjVu3",
  "dispatchId": "DcBMPCvx3nSWV4aVkivJQA",
  "containerId": "YaWsqwD5Qqz3SLcqikB6ej",
  "customerId": "7M6wvxG2KZRz7DMqmpZD5Y",
  "customerName": "PIONEER",
  "customerEmail": "bigredtags@yahoo.com",
  "customerPhone": "6268311082",
  "customerDeliveryInfo": [
    {
      "city": "City of Industry",
      "consignee": "EP",
      "consigneePhone": "6268311082",
      "state": "CA",
      "street": "1521 PIER AVE"
    },
    {
      "city": "Santa Fe Springs",
      "consignee": "COSMIC",
      "consigneePhone": "6268311082",
      "state": "CA",
      "street": "12062 VALLEY VIEW ST. SUITE 204 B"
    }
  ],
  "dispatcherId": "b8AnejLc3ziDT4U4fRb5dk",
  "dispatcherName": "skyler-yang",
  "dispatcherEmail": "skyler172789399@gmail.com",
  "driverId": "RE3Ndtp5sbtiuDy3h3hfUk",
  "driverName": "TIM ZHANG",
  "driverEmail": "dhuang309@hotmail.com",
  "driverPhone": "6268311082",
  "driverLicense": "B1234567",
  "driverLicensePlate": "",
  "containerNO": "LOAD1111111",
  "containerLoad": "loaded",
  "containerSize": "40HQ",
  "containerCarrier": "ACLU",
  "containerFromState": "CA",
  "containerFromCity": "Los Angeles",
  "containerFromStreet": "one",
  "containerFromContact": "",
  "containerFromContactPhone": "",
  "containerToState": "CA",
  "containerToCity": "Chino",
  "containerToStreet": "336699  AABBCC BLVD",
  "containerToContact": "",
  "containerToContactPhone": "",
  "containerChassisType": 1,
  "containerChassisOwner": "",
  "containerChassisSizes": ["40"],
  "containerChassisNumber": "itik203",
  "containerEtaAt": 1742972400000,
  "containerLfdAt": 1743231600000,
  "containerLrdAt": 1743145200000,
  "containerEmptyAt": 0,
  "containerMblNO": "",
  "containerSealNO": "",
  "containerVoyage": "",
  "containerHouseNO": "",
  "containerAdditionalInfo": [],
  "containerTagInfo": {
    "ctf": false,
    "deliveryOnHold": false,
    "forStorage": false,
    "hazmat": false,
    "holdCTF": false,
    "holdCustom": false,
    "holdFreight": false,
    "holdPerpass": false,
    "hot": false,
    "link": "",
    "live": false,
    "onVessel": false,
    "overweight": false,
    "returnLocation": "",
    "socShipperOwnedContainer": false,
    "tmfPierPass": false,
    "weight": 0,
    "weightUnit": "LB"
  },
  "containerNotes":
    "must deliver by 3pm, if not call dispatcher,must deliver by 3pm, if not call dispatcher,must deliver by 3pm, if not call dispatcher,must deliver by 3pm, if not call dispatcher",
  "containerCreatedUserId": "b8AnejLc3ziDT4U4fRb5dk",
  "containerCreatedUserName": "skyler-yang",
  "containerCreatedAt": 1742793981089,
  "dispatchPinNo": "1233",
  "dispatchFromState": "CA",
  "dispatchFromCity": "Los Angeles",
  "dispatchFromStreet": "one",
  "dispatchFromContact": "",
  "dispatchFromContactPhone": "",
  "dispatchToState": "CA",
  "dispatchToCity": "Chino",
  "dispatchToStreet": "336699  AABBCC BLVD",
  "dispatchToContact": "",
  "dispatchToContactPhone": "",
  "dispatchChassisType": 0,
  "dispatchChassisOwner": "",
  "dispatchChassisSizes": ["40"],
  "dispatchChassisNumber": "itik203",
  "dispatchDriverUploadChassisType": 0,
  "dispatchDriverUploadChassisOwner": "",
  "dispatchDriverUploadChassisSizes": ["40"],
  "dispatchDriverUploadChassisNumber": "itik203",
  "dispatchApptStartAt": 1742828400000,
  "dispatchApptEndAt": 1742832000000,
  "dispatchLoadOutImages": [
    {
      "contentType": "image/jpeg",
      "fileId": "0c69d995632b4896967feefb66313497",
      "filePath": "/dispatch/YaWsqwD5Qqz3SLcqikB6ej/DcBMPCvx3nSWV4aVkivJQA/0c69d995632b4896967feefb66313497.jpg",
      "group": "ibwiry7jtpyfy4jmezjvu3",
      "name": "image.jpg",
      "size": 752879,
      "smallFilePath": "/dispatch/YaWsqwD5Qqz3SLcqikB6ej/DcBMPCvx3nSWV4aVkivJQA/0c69d995632b4896967feefb66313497-small.jpg"
    },
    {
      "contentType": "image/jpeg",
      "fileId": "8c801472c00349c1a860463e73e7fdad",
      "filePath": "/dispatch/YaWsqwD5Qqz3SLcqikB6ej/DcBMPCvx3nSWV4aVkivJQA/8c801472c00349c1a860463e73e7fdad.jpg",
      "group": "ibwiry7jtpyfy4jmezjvu3",
      "name": "image.jpg",
      "size": 446851,
      "smallFilePath": "/dispatch/YaWsqwD5Qqz3SLcqikB6ej/DcBMPCvx3nSWV4aVkivJQA/8c801472c00349c1a860463e73e7fdad-small.jpg"
    }
  ],
  "dispatchLoadOutAt": 1742799966142,
  "dispatchEmptyInImages": [],
  "dispatchEmptyInAt": 0,
  "dispatchDeliveredImages": [
    {
      "contentType": "image/jpeg",
      "fileId": "77825d814b4749f8bdb14d84009e1302",
      "filePath": "/dispatch/YaWsqwD5Qqz3SLcqikB6ej/DcBMPCvx3nSWV4aVkivJQA/77825d814b4749f8bdb14d84009e1302.jpg",
      "group": "ibwiry7jtpyfy4jmezjvu3",
      "name": "image.jpg",
      "size": 919080,
      "smallFilePath": "/dispatch/YaWsqwD5Qqz3SLcqikB6ej/DcBMPCvx3nSWV4aVkivJQA/77825d814b4749f8bdb14d84009e1302-small.jpg"
    }
  ],
  "dispatchDeliveredAt": 1742799966142,
  "dispatchStatus": 1,
  "dispatchNotes": "",
  "dispatchAt": 1742795574527,
  "dispatchCompletedUserId": "b8AnejLc3ziDT4U4fRb5dk",
  "dispatchCompletedUserName": "skyler-yang",
  "dispatchCompletedAt": 1742800224517
})

const newData = ref({
  "dbId": 39,
  "orgId": "ibwirY7jtPyFY4JmEZjVu3",
  "dispatchId": "Ww33SXfpz2RbCUi8NnJgj6",
  "containerId": "c3SqeBaEKNmdgpPi3TcbuX",
  "customerId": "7M6wvxG2KZRz7DMqmpZD5Y",
  "customerName": "PIONEER",
  "customerEmail": "bigredtags@yahoo.com",
  "customerPhone": "6268311082",
  "customerDeliveryInfo": [
    {
      "city": "City of Industry",
      "consignee": "EP",
      "consigneePhone": "6268311082",
      "state": "CA",
      "street": "1521 PIER AVE"
    },
    {
      "city": "Santa Fe Springs",
      "consignee": "COSMIC",
      "consigneePhone": "6268311082",
      "state": "CA",
      "street": "12062 VALLEY VIEW ST. SUITE 204 B"
    }
  ],
  "dispatcherId": "b8AnejLc3ziDT4U4fRb5dk",
  "dispatcherName": "skyler-yang",
  "dispatcherEmail": "skyler172789399@gmail.com",
  "dispatcherPhone": "6262678010",
  "driverId": "RE3Ndtp5sbtiuDy3h3hfUk",
  "driverName": "TIM ZHANG",
  "driverEmail": "dhuang309@hotmail.com",
  "driverPhone": "6268311082",
  "driverLicense": "B1234567",
  "driverLicensePlate": "",
  "containerNO": "EMPTY111111",
  "containerLoad": "empty",
  "containerSize": "40HQ",
  "containerCarrier": "ACLU",
  "containerFromState": "CA",
  "containerFromCity": "Chino",
  "containerFromStreet": "336699  AABBCC BLVD",
  "containerFromContact": "",
  "containerFromContactPhone": "",
  "containerToState": "CA",
  "containerToCity": "Los Angeles",
  "containerToStreet": "one",
  "containerToContact": "",
  "containerToContactPhone": "",
  "containerChassisType": 1,
  "containerChassisOwner": "",
  "containerChassisSizes": ["40"],
  "containerChassisNumber": "ITIK203",
  "containerEtaAt": 1742972400000,
  "containerLfdAt": 1743145200000,
  "containerLrdAt": 1743145200000,
  "containerEmptyAt": 1742721300000,
  "containerMblNO": "",
  "containerSealNO": "",
  "containerVoyage": "",
  "containerHouseNO": "",
  "containerAdditionalInfo": [],
  "containerTagInfo": {
    "ctf": false,
    "deliveryOnHold": false,
    "forStorage": false,
    "hazmat": false,
    "holdCTF": false,
    "holdCustom": false,
    "holdFreight": false,
    "holdPerpass": false,
    "hot": false,
    "link": "",
    "live": false,
    "onVessel": false,
    "overweight": false,
    "returnLocation": "",
    "socShipperOwnedContainer": false,
    "tmfPierPass": false,
    "weight": 0,
    "weightUnit": "LB"
  },
  "containerNotes": "gate code 2556",
  "containerCreatedUserId": "b8AnejLc3ziDT4U4fRb5dk",
  "containerCreatedUserName": "skyler-yang",
  "containerCreatedAt": 1742794011532,
  "dispatchPinNo": "1233",
  "dispatchFromState": "CA",
  "dispatchFromCity": "Chino",
  "dispatchFromStreet": "336699  AABBCC BLVD",
  "dispatchFromContact": "",
  "dispatchFromContactPhone": "",
  "dispatchToState": "CA",
  "dispatchToCity": "Los Angeles",
  "dispatchToStreet": "one",
  "dispatchToContact": "",
  "dispatchToContactPhone": "",
  "dispatchChassisType": 0,
  "dispatchChassisOwner": "",
  "dispatchChassisSizes": ["40"],
  "dispatchChassisNumber": "ITIK203",
  "dispatchDriverUploadChassisType": 0,
  "dispatchDriverUploadChassisOwner": "",
  "dispatchDriverUploadChassisSizes": ["40", "45"],
  "dispatchDriverUploadChassisNumber": "ITIK203",
  "dispatchApptStartAt": 1742799600000,
  "dispatchApptEndAt": 1742803200000,
  "dispatchLoadOutImages": [
    {
      "contentType": "image/jpeg",
      "fileId": "8e07a45a936f49ea80f4372e105d4041",
      "filePath": "/dispatch/c3SqeBaEKNmdgpPi3TcbuX/Ww33SXfpz2RbCUi8NnJgj6/8e07a45a936f49ea80f4372e105d4041.jpg",
      "group": "ibwiry7jtpyfy4jmezjvu3",
      "name": "image.jpg",
      "size": 895976,
      "smallFilePath": "/dispatch/c3SqeBaEKNmdgpPi3TcbuX/Ww33SXfpz2RbCUi8NnJgj6/8e07a45a936f49ea80f4372e105d4041-small.jpg"
    }
  ],
  "dispatchLoadOutAt": 1742796579686,
  "dispatchEmptyInImages": [
    {
      "contentType": "image/jpeg",
      "fileId": "49a54cf38763455d90cdb8f0d1e57fd8",
      "filePath": "/dispatch/c3SqeBaEKNmdgpPi3TcbuX/Ww33SXfpz2RbCUi8NnJgj6/49a54cf38763455d90cdb8f0d1e57fd8.jpg",
      "group": "ibwiry7jtpyfy4jmezjvu3",
      "name": "image.jpg",
      "size": 892944,
      "smallFilePath": "/dispatch/c3SqeBaEKNmdgpPi3TcbuX/Ww33SXfpz2RbCUi8NnJgj6/49a54cf38763455d90cdb8f0d1e57fd8-small.jpg"
    }
  ],
  "dispatchEmptyInAt": 1742796579686,
  "dispatchDeliveredImages": [
    {
      "contentType": "image/jpeg",
      "fileId": "af028e9bb20b43e888e3c67e46b0577b",
      "filePath": "/dispatch/c3SqeBaEKNmdgpPi3TcbuX/Ww33SXfpz2RbCUi8NnJgj6/af028e9bb20b43e888e3c67e46b0577b.jpg",
      "group": "ibwiry7jtpyfy4jmezjvu3",
      "name": "image.jpg",
      "size": 904013,
      "smallFilePath": "/dispatch/c3SqeBaEKNmdgpPi3TcbuX/Ww33SXfpz2RbCUi8NnJgj6/af028e9bb20b43e888e3c67e46b0577b-small.jpg"
    }
  ],
  "dispatchDeliveredAt": 1742796579686,
  "dispatchStatus": 1,
  "dispatchNotes": "",
  "dispatchAt": 1742795574527,
  "dispatchCompletedUserId": "b8AnejLc3ziDT4U4fRb5dk",
  "dispatchCompletedUserName": "skyler-yang",
  "dispatchCompletedAt": 1742796973131
})

const isLoading = ref(false)

const queryInfo = ref<QueryInfo>(new QueryInfo({}, ["keywords"], { index: 1, size: 20 }, { field: "createdAt", order: "desc" }, ["createdAt"]))

const tableList = ref({
  total: 0,
  data: [] as Array<any>
})

const setRowClass = ({ row }) => {
  return !row.oldContent && !row.newContent ? "hide-expand-icon" : ""
}

function search() {
  const list: Array<any> = []
  for (let i = 0; i < 20; i++) {
    list.push({
      activityId: i,
      hasContent: i % 2 === 0,
      oldContent: i % 2 === 0 ? oldData.value : null,
      newContent: i % 2 === 0 ? newData.value : null,
      description: "Update application",
      user: "skyler",
      createdAt: new Date().getTime()
    })
  }
  tableList.value.data = list
  tableList.value.total = tableList.value.data.length * 3
}

onMounted(() => {
  search()
})
</script>

<template>
  <div class="d-flex gap-3">
    <div class="header-icon">
      <el-icon :size="18">
        <IconMdiAccountFileText />
      </el-icon>
    </div>
    <el-text class="f-bold" size="large">{{ t("label.activities") }}</el-text>
  </div>

  <v-table
    v-loading="isLoading"
    v-model:page-info="queryInfo.pageInfo"
    :data="tableList.data"
    :max-height="tableHeight"
    :row-class-name="setRowClass"
    :total="tableList.total"
    class="mt-5"
    row-key="activityId"
    @page-change="search">
    <el-table-column align="left" class-name="expand-column" type="expand" width="24">
      <template #default="scope">
        <div class="px-5">
          <v-monaco-diff-editor
            v-if="scope.row.hasContent"
            :new-data="JSON.stringify(scope.row.oldContent, null, 4)"
            :old-data="JSON.stringify(scope.row.newContent, null, 4)"
            language="json" />
        </div>
      </template>
    </el-table-column>
    <el-table-column :label="t('label.description')" min-width="200" prop="description" />
    <el-table-column :label="t('label.operator')" min-width="140" prop="user" />
    <el-table-column :label="t('label.date')" width="140">
      <template #default="scope">
        <v-date-view :timestamp="scope.row.createdAt" />
      </template>
    </el-table-column>
  </v-table>
</template>

<style lang="scss" scoped>
.header-icon {
  color: var(--el-color-primary);
  background-color: var(--el-color-primary-light-9);
  border-radius: 50%;
  padding: 8px;
  width: 18px;
  height: 18px;
  display: flex;
  align-items: center;
  justify-content: center;
}

:deep(.expand-column) {
  .cell {
    padding: 0 4px 0 8px;
  }
}

:deep(.hide-expand-icon) {
  .expand-column .cell {
    padding-top: 4px;
    display: none;
  }
}
</style>
