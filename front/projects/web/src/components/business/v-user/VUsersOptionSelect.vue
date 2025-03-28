<script lang="ts" setup>
import { UserInfo } from "@/types"
import { map } from "lodash-es"

const props = withDefaults(
  defineProps<{
    options?: UserInfo[]
    filterable?: boolean
    multiple?: boolean
    placeholder?: string
  }>(),
  {
    filterable: true,
    multiple: true,
    placeholder: ""
  }
)

const { t } = useI18n()
const userStore = useUserStore()

const isLoading = ref(false)
const users = defineModel<string[]>()
const userList = ref<UserInfo[]>([])
const selectOptions = computed(() => map(props.options || userList.value, x => ({ label: x.username, value: x.userId })))

onMounted(() => {
  if (!props.options) {
    isLoading.value = true
    userService
      .list()
      .then(res => {
        userList.value = res
      })
      .finally(() => (isLoading.value = false))
  }
})
</script>

<template>
  <el-select-v2
    v-model="users"
    :filterable="props.filterable"
    :loading-text="t('message.loading')"
    :multiple="props.multiple"
    :options="selectOptions"
    :placeholder="props.placeholder">
    <template v-if="isLoading" #prefix>
      <el-button :loading="isLoading" link />
    </template>
    <template v-if="userStore.isAdmin || userStore.isSupperAdmin" #footer>
      <div class="text-align-right">
        <el-link href="/ws/user-related/users" target="_blank" type="primary">
          <strong>{{ t("btn.goToAddUser") }}</strong>
        </el-link>
      </div>
    </template>
  </el-select-v2>
</template>

<style lang="scss" scoped></style>
