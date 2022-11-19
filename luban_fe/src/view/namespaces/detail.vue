<template>
    <div>
      <el-row :gutter="20" class="row-box">
        <el-col :span="24">
          <el-card class="el-card">
            <detail-basic :item="form" />
          </el-card>
        </el-col>
      </el-row>
    </div>
</template>

<script>
  import DetailBasic from "@/components/detail/detail-basic"
  import { formatTimeToStr } from '@/utils/date'
    export default {
        name: 'DetailBlock',
      components: { DetailBasic},
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
                activeName: "pods",
                ContainerTitle: '',
                dialogVisibleContainerlog: false,
                currentValuedetailurl: '',
                terminal: {
                  pid: 1,
                  name: 'terminal',
                  cols: 150,
                  rows: 23
                },
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
        },
        methods: {
            handleClick (tab) {
                this.activeName = tab.name
            },
        }
    }
</script>
<style scoped lang="scss">

</style>
