<script lang="ts" setup>
const { t } = useI18n()

// function search() {}

const searchForm = ref({
  name: "",
  email: "",
  phone: "",
  group: ""
})

const tableData = ref([
  { name: "John Doe", email: "john@example.com", phone: "123-456-7890", description: "Developer", group: "Admin" },
  { name: "Jane Smith", email: "jane@example.com", phone: "987-654-3210", description: "Designer", group: "User" }
])

const filteredData = ref([...tableData.value])

function search() {
  filteredData.value = tableData.value.filter(item => {
    return (
      (!searchForm.value.name || item.name.includes(searchForm.value.name)) &&
      (!searchForm.value.email || item.email.includes(searchForm.value.email)) &&
      (!searchForm.value.phone || item.phone.includes(searchForm.value.phone)) &&
      (!searchForm.value.group || item.group.includes(searchForm.value.group))
    )
  })
}
</script>

<template>
  <v-card>
    <!--    <el-form @submit.prevent="search()">-->
    <!--      <el-form-item>-->
    <!--        <v-input v-model="searchForm.name" placeholder="Name" />-->
    <!--      </el-form-item>-->
    <!--    </el-form>-->

    <el-form @submit.prevent="search">
      <el-form-item>
        <v-input v-model="searchForm.name" placeholder="Name" />
      </el-form-item>
      <el-form-item>
        <v-input v-model="searchForm.email" placeholder="Email" />
      </el-form-item>
      <el-form-item>
        <v-input v-model="searchForm.phone" placeholder="Phone" />
      </el-form-item>
      <el-form-item>
        <v-input v-model="searchForm.group" placeholder="Group" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="search">{{ t("Search") }}</el-button>
      </el-form-item>
    </el-form>
    <v-table :data="filteredData">
      <el-table-column label="Name" prop="name" sortable />
      <el-table-column label="Email" prop="email" />
      <el-table-column label="Phone" prop="phone" />
      <el-table-column label="Description" prop="description" />
      <el-table-column label="Group" prop="group" />
      <el-table-column label="Action">
        <template #default="scope">
          <el-button type="text" @click="() => handleEdit(scope.row)">Edit</el-button>
          <el-button type="text" @click="() => handleDelete(scope.row)">Delete</el-button>
        </template>
      </el-table-column>
    </v-table>
  </v-card>
</template>

<style lang="scss" scoped></style>
