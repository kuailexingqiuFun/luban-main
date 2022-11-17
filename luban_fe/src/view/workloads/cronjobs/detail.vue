<template>
  <div>
    <el-row :gutter="20" class="row-box">
      <el-col :span="24">
        <el-card class="el-card">
          <detail-basic :item="form" />
        </el-card>
      </el-col>
      <el-col :span="24">
        <el-tabs v-model="activeName" tab-position="top" type="border-card"
                 @tab-click="handleClick" ref=tabs>
          <el-tab-pane label="Jobs" name="Jobs">
            <el-table :data="jobs">
              <el-table-column label="状态" prop="status.succeeded" min-width="30">
                <template v-slot:default="{row}">
                  <el-button v-if="row.status.succeeded" type="success" size="mini" plain round>
                    Succeeded
                  </el-button>
                  <el-button v-if="row.status.failed" type="warning" size="mini" plain round>
                    Failed
                  </el-button>
                  <el-button v-if="row.status.active" type="success" size="mini" plain round>
                    Active
                  </el-button>
                </template>
              </el-table-column>
              <el-table-column label="名称" prop="metadata.name" min-width="90" show-overflow-tooltip>
                <template v-slot:default="{row}">
                  <span class="span-link">{{
                      row.metadata.name
                    }}
                  </span>
                </template>
              </el-table-column>

              <el-table-column label="需完成数" min-width="30">
                <template v-slot:default="{row}">
                  {{ row.spec.completions }}
                </template>
              </el-table-column>
              <el-table-column label="完成状态" min-width="60">
                <template v-slot:default="{row}">
                  <el-tag style="margin-left: 5px" v-if="row.status.active" type="info">active: *{{row.status.active}}</el-tag>
                  <el-tag style="margin-left: 5px" v-if="row.status.succeeded" type="success">succeeded: {{row.status.succeeded}}</el-tag>
                  <el-tag style="margin-left: 5px" v-if="row.status.failed" type="danger">failed: {{row.status.failed}}</el-tag>
                </template>
              </el-table-column>
              <el-table-column label="耗时" min-width="30">
                <template v-slot:default="{row}">
                  {{ getDuration(row) }}
                </template>
              </el-table-column>
              <el-table-column label="存活时间" min-width="60" prop="metadata.creationTimestamp" fix>
                <template v-slot:default="{row}">
                  {{ row.metadata.creationTimestamp | formatDate }}
                </template>
              </el-table-column>
            </el-table>
          </el-tab-pane>
          <el-tab-pane name="conditions" label="运行时信息">
            <detail-conditions :conditions="form.status.conditions" />
          </el-tab-pane>
        </el-tabs>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import DetailBasic from "@/components/detail/detail-basic"
import DetailConditions from "@/components/detail/detail-conditions"
import { formatTimeToStr } from '@/utils/date'
import { JobsList} from '@/api/kubernetes/workloads/jobs'
export default {
  name: 'DetailBlock',
  components: { DetailBasic, DetailConditions},
  props: {
    cluster_id: Number,
    form: {
      type: Object,
      default: function() {
        return {}
      }
    }
  },
  data() {
    return {
      name: '',
      activeName: "Jobs",
      ContainerTitle: '',
      dialogVisibleContainerlog: false,
      currentValuedetailurl: '',
      terminal: {
        pid: 1,
        name: 'terminal',
        cols: 150,
        rows: 23
      },
      jobs: [],
      selectors: "",
    }
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time !== '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    }
  },
  created() {
    this.getDetail()
  },
  methods: {
    getDetail() {
      this.loading = true
      JobsList(this.cluster_id, "", "","", this.form.metadata.namespace,"").then((res) => {
        this.jobs = []
        for (const job of res.data.items) {
          if (job.metadata.name.indexOf("-") !== -1) {
            if (job.metadata.name.substring(0, job.metadata.name.lastIndexOf("-")) === this.form.metadata.name) {
              this.jobs.push(job)
            }
          }
        }
        this.loading = false
      })
    },
    getDuration(row) {
      let startTime = new Date(row.status.startTime)
      if (!row.status.completionTime) {
        return "/"
      } else {
        let endTime = new Date(row.status.completionTime)
        let t = Math.floor((endTime - startTime) / 1000)
        if (t % 60 !== 0) {
          return (t % 60) + " mins"
        }
        if (t % 3600 !== 0) {
          return (t % 60) + " hours ago"
        }
        return Math.floor((endTime - startTime) / 1000) + "S"
      }
    },
    handleClick (tab) {
      this.activeName = tab.name
    },
  }
}
</script>
<style scoped lang="scss">

</style>
