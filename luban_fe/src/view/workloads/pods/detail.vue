<template>
  <!--应用详情-->
  <!--最外层模态框-->
  <div>
    <div class="console-sub-title custom-sub-title top-sub clearfix">
      <div class="pull-left">
        <h4>基本信息</h4>
      </div>
    </div>
    <table
        class="table-default-viewer current-row"
        style="width: 100%">
      <tbody>
      <tr>
        <td style="width: 50%">
          <span>名称</span>
          <span class="margin-right"> :</span>
          <span> {{value.metadata.name}} </span>
        </td>
        <td style="width: 50%">
          <span>命名空间</span>
          <span class="margin-right"> :</span>
          <span> {{value.metadata.namespace}}</span>
        </td>
      </tr>
      <tr>
        <td style="width: 50%">
          <span>状态</span>
          <span class="margin-right"> :</span>
          <span> {{value.status.phase}} </span>
        </td>
        <td style="width: 50%">
          <span>创建时间</span>
          <span class="margin-right"> :</span>
          <span> {{value.metadata.creationTimestamp |formatDate}}</span>
        </td>
      </tr>
      <tr>
        <td style="width: 50%">
          <span>节点</span>
          <span class="margin-right"> :</span>
          <span> {{value.status.hostIP}} </span>
        </td>
        <td style="width: 50%">
          <span>Pod IP</span>
          <span class="margin-right"> :</span>
          <span> {{value.status.podIP}}</span>
        </td>
      </tr>
      <tr>
        <td style="width: 50%">
          <span>标签</span>
          <span class="margin-right"> :</span>
          <span v-for="(k, v, index) in value.metadata.labels" :key="index">
                          <span class="label-custom" type="info">{{ k }}: {{ v }}</span>
                     </span>
        </td>
        <td style="width: 50%">
          <span>注解</span>
          <span> :</span>
          <span v-for="(k, v, index) in value.metadata.annotations" :key="index">
                          <span class="margin-right label-custom" type="info">{{ k }}: {{ v }}</span>
                     </span>
        </td>
      </tr>
      </tbody>
    </table>
    <div class="console-sub-title custom-sub-title top-sub clearfix">
      <div class="pull-left">
        <h4>现状详情</h4>
      </div>
    </div>
    <div>
      <el-table
          :data="value.status.conditions"
          stripe
          style="width: 100%">
        <el-table-column
            label="类型"
            prop="type"
            type="scope">
          <template slot-scope="scope" >
            <span>{{ scope.row.type }} </span>
          </template>
        </el-table-column>
        <el-table-column
            label="状态"
            prop="status"
            type="scope">
          <template slot-scope="scope" >
            <span>{{ scope.row.status }} </span>
          </template>
        </el-table-column>
        <el-table-column
            label="更新时间"
            prop="status"
            type="scope">
          <template slot-scope="scope" >
            <span>{{ scope.row.lastTransitionTime |formatDate }} </span>
          </template>
        </el-table-column>
        <el-table-column
            label="内容"
            prop="reason"
            type="scope">
          <template slot-scope="scope" >
            <span v-if="scope.row.reason">{{ scope.row.reason }} </span>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column
            label="消息"
            prop="message"
            type="scope">
          <template slot-scope="scope" >
            <span v-if="scope.row.message">{{scope.row.message}}</span>
            <span v-else>-</span>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="container">
      <el-tabs style="width:100%;">
        <el-tab-pane label="容器">
          <div style="width: 100%">
            <el-table
                :data="value.spec.containers"
                style="width: 100%">
              <el-table-column type="expand">
                <template slot-scope="props">
                  <el-form label-position="" inline class="demo-table-expand">
                    <el-form-item label="镜像拉取策略: ">
                      <span>{{ props.row.imagePullPolicy }} </span>
                    </el-form-item>
                    <el-form-item label="环境变量: ">
                                 <span v-for="(e, index) in props.row.env" :key="index">
                                   <span class="margin-right label-custom" type="info">{{ e.name}}: {{ e.value }}</span>
                                </span>
                    </el-form-item>
                    <el-form-item label="命令: ">
                      <span>{{ props.row.command }}</span>
                    </el-form-item>
                    <el-form-item v-if="props.row.resources.requests" label="所需资源: ">
                      <span>CPU: {{ props.row.resources.requests.cpu}}  </span>
                      <span>Memory:  {{ props.row.resources.requests.memory}} </span>
                    </el-form-item>
                    <el-form-item v-if="props.row.resources.requests" label="资源限制: ">
                      <span>CPU: {{ props.row.resources.limits.cpu}}  </span>
                      <span>Memory:  {{ props.row.resources.limits.memory}} </span>
                    </el-form-item>
                  </el-form>
                </template>
              </el-table-column>
              <el-table-column
                  label="名称"
                  prop="name"
                  type="scope">
                <template slot-scope="scope" >
                  <span>{{ scope.row.name }} </span>
                </template>
              </el-table-column>
              <el-table-column
                  label="镜像"
                  prop="image"
                  type="scope">
                <template slot-scope="scope" >
                  <span>{{ scope.row.image }} </span>
                </template>
              </el-table-column>
              <el-table-column
                  label="端口"
                  prop="ports"
                  type="scope">
                <template slot-scope="scope" >
                            <span v-for="(p, index) in scope.row.ports" :key="index">
                               <span class="margin-right label-custom" type="info">{{ p.protocol }}: {{ p.containerPort }}</span>
                            </span>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-tab-pane>
        <el-tab-pane label="事件" >
          <el-table
              :data="value.events"
              stripe
              style="width: 100%">
            <el-table-column
                label="类型"
                prop="type"
                type="scope">
              <template slot-scope="scope" >
                <span>{{ scope.row.type }} </span>
              </template>
            </el-table-column>
            <el-table-column
                label="对象"
                prop="involvedObject"
                type="scope">
              <template slot-scope="scope" >
                <span>{{ scope.row.involvedObject.kind}} : {{ scope.row.involvedObject.name}} </span>
              </template>
            </el-table-column>
            <el-table-column
                label="信息"
                prop="message"
                type="scope">
              <template slot-scope="scope" >
                <span>{{scope.row.message}}</span>
              </template>
            </el-table-column>
            <el-table-column
                label="内容"
                prop="reason"
                type="scope">
              <template slot-scope="scope" >
                <span>{{ scope.row.reason }} </span>
              </template>
            </el-table-column>
            <el-table-column
                label="时间"
                prop="eventTime"
                type="scope">
              <template slot-scope="scope" >
                <span>{{scope.row.lastTimestamp |formatDate}}</span>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
        <el-tab-pane label="创建者" >
          <el-table
              :data="value.metadata.ownerReferences"
              stripe
              style="width: 100%">
            <el-table-column
                label="名称"
                prop="name"
                type="scope">
              <template slot-scope="scope" >
                <span>{{ scope.row.name }} </span>
              </template>
            </el-table-column>
            <el-table-column
                label="类型"
                prop="type"
                type="scope">
              <template slot-scope="scope" >
                <span>{{ scope.row.kind }} </span>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
        <el-tab-pane label="初始化容器" >
          <div style="width: 100%">
            <el-table
                :data="value.spec.initContainers"
                style="width: 100%">
              <el-table-column type="expand">
                <template slot-scope="props">
                  <el-form label-position="" inline class="demo-table-expand">
                    <el-form-item label="镜像拉取策略: ">
                      <span>{{ props.row.imagePullPolicy }} </span>
                    </el-form-item>
                    <el-form-item label="环境变量: ">
                                 <span v-for="(e, index) in props.row.env" :key="index">
                                   <span class="margin-right label-custom" type="info">{{ e.name}}: {{ e.value }}</span>
                                </span>
                    </el-form-item>
                    <el-form-item label="命令: ">
                      <span>{{ props.row.command }}</span>
                    </el-form-item>
                    <el-form-item v-if="props.row.resources" label="所需资源: ">
                      <span v-if="rops.row.resources.requests.cpu">CPU: {{ props.row.resources.requests.cpu}}  </span>
                      <span>Memory:  {{ props.row.resources.requests.memory}} </span>
                    </el-form-item>
                    <el-form-item  v-if="props.row.resources" label="资源限制: ">
                      <span>CPU: {{ props.row.resources.limits.cpu}}  </span>
                      <span>Memory:  {{ props.row.resources.limits.memory}} </span>
                    </el-form-item>
                  </el-form>
                </template>
              </el-table-column>
              <el-table-column
                  label="名称"
                  prop="name"
                  type="scope">
                <template slot-scope="scope" >
                  <span>{{ scope.row.name }} </span>
                </template>
              </el-table-column>
              <el-table-column
                  label="镜像"
                  prop="image"
                  type="scope">
                <template slot-scope="scope" >
                  <span>{{ scope.row.image }} </span>
                </template>
              </el-table-column>
              <el-table-column
                  label="端口"
                  prop="ports"
                  type="scope">
                <template slot-scope="scope" >
                            <span v-for="(p, index) in scope.row.ports" :key="index">
                               <span class="margin-right label-custom" type="info">{{ p.protocol }}: {{ p.containerPort }}</span>
                            </span>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-tab-pane>
        <el-tab-pane label="存储" >
          <el-table
              v-if="value.spec.volumes"
              :data="value.spec.volumes"
              stripe
              style="width: 100%">
            <el-table-column
                label="名称"
                prop="name"
                type="scope">
              <template slot-scope="scope" >
                <span>{{ scope.row.name }} </span>
              </template>
            </el-table-column>
            <el-table-column
                label="类型"
                prop="type"
                type="scope">
              <template slot-scope="scope" >
                <span v-if="scope.row.hostPath">hostPath</span>
                <span v-else-if="scope.row.configMap">configMap</span>
                <span v-else-if="scope.row.secret">secret</span>
                <span v-else-if="scope.row.nfs">nfs</span>
                <span v-else-if="scope.row.glusterfs">glusterfs</span>
                <span v-else-if="scope.row.emptyDir">emptyDir</span>
                <span v-else-if="scope.row.cephfs">cephfs</span>
                <span v-else>-</span>
              </template>
            </el-table-column>
            <el-table-column
                label="详情"
                prop="detail"
                type="scope">
              <template slot-scope="scope" >
                <div v-if="scope.row.hostPath">
                  <div v-for="(path_k, path_v, path_i) in scope.row.hostPath" :key="path_i"> {{ path_v }}: {{ path_k }}</div>
                </div>
                <div v-else-if="scope.row.configMap">
                  <div v-for="(path_k, path_v, path_i) in scope.row.configMap" :key="path_i"> {{ path_v }}: {{ path_k }}</div>
                </div>
                <div v-else-if="scope.row.secret">
                  <div v-for="(path_k, path_v, path_i) in scope.row.secret" :key="path_i"> {{ path_v }}: {{ path_k }}</div>
                </div>
                <div v-else-if="scope.row.nfs">
                  <div v-for="(path_k, path_v, path_i) in scope.row.nfs" :key="path_i"> {{ path_v }}: {{ path_k }}</div>
                </div>
                <div v-else-if="scope.row.glusterfs">
                  <div v-for="(path_k, path_v, path_i) in scope.row.glusterfs" :key="path_i"> {{ path_v }}: {{ path_k }}</div>
                </div>
                <div v-else-if="scope.row.emptyDir">
                  <div v-for="(path_k, path_v, path_i) in scope.row.emptyDir" :key="path_i"> {{ path_v }}: {{ path_k }}</div>
                </div>
                <div v-else-if="scope.row.cephfs">
                  <div v-for="(path_k, path_v, path_i) in scope.row.cephfs" :key="path_i"> {{ path_v }}: {{ path_k }}</div>
                </div>
                <span v-else>-</span>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script>
import { formatTimeToStr } from '@/utils/date'
export default {
  name: 'DetailBlock',
  props: {
    value: {
      type: Object,
      default: function() {
        return {}
      }
    }
  },
  data() {
    return {
      name: '',
      ContainerTitle: '',
      dialogVisibleContainerlog: false,
      currentValuedetailurl: '',
      terminal: {
        pid: 1,
        name: 'terminal',
        cols: 150,
        rows: 23
      },
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
  methods: {
  }
}
</script>
<style scoped lang="scss">

</style>
