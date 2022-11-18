<template>
  <div>
    <el-table ref="multipleTable" :data="value" style="width: 100%" tooltip-effect="dark" row-key="ID">
      <el-table-column type="selection" width="55" />
      <el-table-column
          prop="name"
          label="名称/IP"
          type="scope">
        <template slot-scope="scope">
          <span class="operate-span"  @click="handleDetail(scope.row)" >{{ scope.row.metadata.name}}</span>
          <el-tooltip placement="top">
            <div slot="content" v-for="(v,k, i) in scope.row.metadata.labels" :key="i">
              <span> {{k}}: {{v}}</span>
            </div>
            <i class="el-icon-price-tag" />
          </el-tooltip>
          <div v-for="(value) in scope.row.status.addresses " :key="value.index">
            <span v-if="value.type ==='InternalIP'">{{ value.address}}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="Internal IP" prop="metadata.name" max-width="30px">
        <template #default="scope">
          <div v-for="(address,index) in scope.row.status.addresses" :key="index">
            <span v-if="address.type === 'InternalIP'">{{ address.address }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="状态" prop="status" type="scope">
        <template #default="scope">
          <el-tag>{{ scope.row | handleNodeStatus}}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="角色" prop="metadata.labels" type="scope">
        <template #default="scope">
          <span>{{ scope.row.metadata.labels | handleNodeRole }}</span>
        </template>
      </el-table-column>
      <el-table-column label="CPU用量" prop="cpuUsage" type="scope">
        <template #default="scope">
          <div>
            <div>
              <span>{{ scope.row.cpuUsagePersent }}%</span>
            </div>
            <div>
              <span>{{ scope.row.cpuUsage | CpuFormatData }} / {{ scope.row.status.allocatable.cpu | CpuFormatData }} 核</span>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="内存用量" prop="memoryUsage" type="scope">
        <template #default="scope">
          <div>
            <div>
              <span>{{ scope.row.memoryUsagePersent }}%</span>
            </div>
            <div>
              <span>{{ scope.row.memoryUsage | MemFormatData}} / {{ scope.row.status.allocatable.memory | MemFormatData }} Gi</span>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="版本" prop="version" type="scope" align="center">
        <template #default="scope">
          <div style="text-align: center">
            <span>{{ scope.row.status.nodeInfo.kubeletVersion }}</span><br>
            <span>{{ scope.row.status.nodeInfo.containerRuntimeVersion }}</span><br>
            <span>{{ scope.row.status.nodeInfo.osImage }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column align="left" label="存活时间" prop="metadata.creationTimestamp" width="180">
        <template #default="scope">
          <span>{{ scope.row.metadata.creationTimestamp | formatDate }}</span>
        </template>
      </el-table-column>
      <el-table-column align="left" label="操作">
        <template #default="scope">
          <el-form :inline="true" class="demo-form-inline">
            <el-form-item>
              <el-button size="small" type="primary" link icon="edit" @click="handleUpdate(scope.row)">编辑</el-button>
            </el-form-item>
            <el-form-item>
              <el-popover v-model="scope.row.visible" placement="top">
                <p>确定要删除吗？</p>
                <div style="text-align: right; margin-top: 8px;">
                  <el-button size="small" type="primary" link @click="scope.row.visible = false">取消</el-button>
                  <el-button type="primary" size="small" @click="handleDelete(scope.row)">确定</el-button>
                </div>
                <template #reference>
                  <el-button type="primary" link icon="delete" size="small" @click="scope.row.visible = true">删除</el-button>
                </template>
              </el-popover>
            </el-form-item>
            <el-form-item>
              <el-dropdown>
                <el-button type="primary" link icon="MoreFilled" size="small">更多</el-button>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item @click.native="handleDetail(scope.row)">详情</el-dropdown-item>
<!--                    <el-dropdown-item @click.native="handleYAML(scope.row)">YAML</el-dropdown-item>-->
<!--                    <el-dropdown-item @click.native="handleDrain(scope.row)">节点排空</el-dropdown-item>-->
<!--                    <el-dropdown-item @click.native="handleSchedule(scope.row)">调度设置</el-dropdown-item>-->
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import { AgeFormat } from '@/utils/age'
import { giMemoryFormat, CpuFormat } from '@/utils/unitConvert'
export default {
  name: 'ListBlock',
  filters: {
    restartCounts: function(restarts) {
      let restart = 0
      if (restarts) {
        for (let i = 0; i < restarts.length; i++) {
          restart += restarts[i].restartCount
        }
      }
      return restart
    },
    CpuFormatData: function (value) {
      return CpuFormat(value)
    },
    MemFormatData: function (value) {
      return giMemoryFormat(value)
    },
    formatDate: function(time) {
      if (time != null && time !== '') {
        return AgeFormat(time)
      } else {
        return ''
      }
    },
    handleNodeStatus: function (value) {
      let status = ''
      for (const condition of value.status.conditions) {
        if (condition.type === 'Ready') {
          if (condition.status === 'True') {
            status = 'Ready'
          }
          break
        }
      }

      if (value.spec.unschedulable) {
        status += ', SchedulingDisabled'
      }
      return status
    },
    handleNodeRole: function (value) {
      let role = ''
      for (const label in value) {
        if (label === 'node-role.kubernetes.io') {
          role = value[label]
        }
      }
      return role
    },
  },
  props: {
    value: {
      type: Array,
      default: function() {
        return []
      }
    }
  },
  data() {
    return {
      name: ''
    }
  },
  methods: {
    handleEdit(value) {
      this.$emit('edit', value)
    },
    handleDetail(value) {
      this.$emit('detail', value)
    },
    handleDelete(value) {
      this.$emit('delete', value)
    },
  }
}
</script>
<style scoped>
.operate-span {
  color: #409EFF;
  cursor: pointer;
  font-weight:bold;
}
</style>
