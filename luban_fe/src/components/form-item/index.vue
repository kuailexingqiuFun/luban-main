<template>
  <div>
    <el-select v-if="itemType==='select'" filterable :clearable="!noClear" style="width: 100%;" v-bind="$attrs" v-on="$listeners">
      <el-option v-for="(item, index) in selections" :key="index" :label="item.label" :value="item.value" />
    </el-select>

    <el-select v-if="itemType==='select2'" filterable :clearable="!noClear" style="width: 100%;" v-bind="$attrs" v-on="$listeners">
      <el-option v-for="(item, index) in selections" :key="index" :label="item" :value="item" />
    </el-select>

    <el-input v-if="itemType==='input'" clearable v-bind="$attrs" v-on="$listeners">
      <template v-if="deviderName" slot="append">{{ deviderName }}</template>
    </el-input>

    <el-input v-if="itemType==='password'" type="password" clearable v-bind="$attrs" v-on="$listeners">
      <template v-if="deviderName" slot="append">{{ deviderName }}</template>
    </el-input>

    <el-input onKeypress="return (/[\d]/.test(String.fromCharCode(event.keyCode)))" v-if="itemType==='number'" v-bind="$attrs" v-on="$listeners">
      <template v-if="deviderName" slot="append">{{ deviderName }}</template>
    </el-input>

    <div v-if="itemType==='checkbox'">
      <el-checkbox-group v-bind="$attrs" v-on="$listeners">
        <el-checkbox v-for="(item, index) in checks" :key="index" :label="item.value">{{ item.label }}</el-checkbox>
      </el-checkbox-group>
    </div>

    <div v-if="itemType==='radio' && radioLayout === 'vertical'" style="display: block;">
      <el-radio-group v-bind="$attrs" v-on="$listeners">
        <el-radio v-for="(item, index) in radios" :key="index" :disabled="item.disabledOption" :label="item.value" style="display: block; line-height: 25px">{{ item.label }}
        </el-radio>
      </el-radio-group>
    </div>
    <div v-if="itemType==='radio' && !radioLayout">
      <el-radio-group v-bind="$attrs" v-on="$listeners">
        <el-radio v-for="(item, index) in radios" :key="index" :disabled="item.disabledOption" :label="item.value">
          {{ item.label }}
        </el-radio>
      </el-radio-group>
    </div>
    <el-input v-if="itemType==='textarea'" style="width:100%;" type="textarea" :rows="1" v-bind="$attrs" v-on="$listeners" :autosize="{maxRows:10}">
      <template v-if="deviderName" slot="append">{{ deviderName }}</template>
    </el-input>
  </div>
</template>

<script>
export default {
  name: "FormItem",
  props: {
    itemType: String, // input, select, radio
    selections: Array, // ????????? select
    noClear: Boolean, // ????????? select ???????????????
    deviderName: String, // ????????????????????????
    radios: Array, // ????????? radio
    radioLayout: String, // radio ???????????????????????? vertical ?????? normal ??????
    checks: Array, // ?????????checkbox
  },
}
</script>

<style>
</style>