
const RequiredRule = {
  required: true,
  trigger: "blur",
  message: "输入错误",
}
const SelectRule = {
  required: true,
  trigger: "change",
  message: "选择错误",
}
// 支持小写英文、数字和-
const CommonNameRule = {
  required: true,
  pattern: /^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$/,
  message: "规则错误",
  trigger: "blur"
}
// 支持数字或%
const NumberPersentRule = {
  required: true,
  pattern: /^[0-9]([0-9]*$|([0-9]*?%$))/,
  message: "非数字",
  trigger: "blur"
}
const NumberRule = {
  required: true,
  trigger: "blur",
  min: 1,
  type: "number",
  message: "数字列表"
}
const UrlRule = {
  required: true,
  pattern: /^(https?):\/\/([a-zA-Z0-9.-]+(:[a-zA-Z0-9.&%$-]+)*@)*((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}|([a-zA-Z0-9-]+\.)*[a-zA-Z0-9-]+\.(com|edu|gov|int|mil|net|org|biz|arpa|info|name|pro|aero|coop|museum|[a-zA-Z]{2}))(:[0-9]+)*(\/($|[a-zA-Z0-9.,?'\\+&%$#=~_-]+))*$/,
  message: "错误的url规则",
  trigger: "blur"
}
export default {
  RequiredRule,
  SelectRule,
  NumberPersentRule,
  NumberRule,
  CommonNameRule,
  UrlRule,
}
