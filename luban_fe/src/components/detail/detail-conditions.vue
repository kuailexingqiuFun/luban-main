<template>
  <div>
    <el-table :data="conditions">
      <el-table-column label="类别" prop="type" min-width="50" show-overflow-tooltip />
      <el-table-column label="状态" prop="status" min-width="30" />
      <el-table-column label="消息" min-width="100" show-overflow-tooltip>
        <template v-slot:default="{row}">
          <span v-if="row.message">[{{ row.reason }} ]: {{ row.message }}</span>
          <span v-if="!row.message">---</span>
        </template>
      </el-table-column>
      <el-table-column label="最后更新时间" min-width="50" prop="lastUpdateTime">
        <template v-slot:default="{row}">
          {{ row.lastUpdateTime |formatDate }}
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
  import { formatTimeToStr } from '@/utils/date'
export default {
  name: "DetailConditions",
  props: {
    conditions: Array,
  },
  filters:{
    formatDate: function(time) {
      if (time != null && time !== '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy/MM/dd hh:mm:ss')
      } else {
        return ''
      }
    }
  }
}
</script>
