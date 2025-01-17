<script lang="ts" setup>
import * as monaco from "monaco-editor"
import EditorWorker from "monaco-editor/esm/vs/editor/editor.worker?worker"
import JsonWorker from "monaco-editor/esm/vs/language/json/json.worker?worker"
import CssWorker from "monaco-editor/esm/vs/language/css/css.worker?worker"
import HtmlWorker from "monaco-editor/esm/vs/language/html/html.worker?worker"
import TSWorker from "monaco-editor/esm/vs/language/typescript/ts.worker?worker"

self.MonacoEnvironment = {
  getWorker(_, label) {
    switch (label) {
      case "json":
        return new JsonWorker()
      case "css":
      case "scss":
      case "less":
        return new CssWorker()
      case "html":
      case "handlebars":
      case "razor":
        return new HtmlWorker()
      case "typescript":
      case "javascript":
        return new TSWorker()
      default:
        return new EditorWorker()
    }
  }
}

monaco.languages.typescript.typescriptDefaults.setEagerModelSync(true)

let monacoCtl: monaco.editor.IStandaloneCodeEditor
const monacoRef = useTemplateRef("monacoRef")

onMounted(() => {
  monacoCtl = monaco.editor.create(monacoRef.value!, {
    value: "",
    language: "text",
    theme: "vs", //官方自带三种主题vs, hc-black, or vs-dark
    selectOnLineNumbers: true, //显示行号
    roundedSelection: false,
    readOnly: false, // 只读
    cursorStyle: "line", //光标样式
    automaticLayout: true, //自动布局
    glyphMargin: true, //字形边缘
    useTabStops: false,
    fontSize: 15, //字体大小
    quickSuggestionsDelay: 10, //代码提示延时
    minimap: {
      enabled: true,
      side: "right"
    },
    scrollBeyondLastLine: false,
    overviewRulerBorder: false,
    formatOnPaste: true
  })
})

onBeforeUnmount(() => {
  if (monacoCtl) {
    monacoCtl.dispose()
  }
})
</script>

<template>
  <div class="monacoBox">
    <div ref="monacoRef" class="monacoContent" />
  </div>
</template>

<style lang="scss" scoped>
.monacoBox {
  width: 100%;
  height: 100%;

  .monacoContent {
    width: 100%;
    height: 100%;
  }

  :deep(.margin-view-overlays) {
    background-color: #7b7979;
  }
}
</style>
