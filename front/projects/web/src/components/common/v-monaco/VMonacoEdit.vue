<script lang="ts" setup>
import screenfull from "screenfull"
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

const boxRef = useTemplateRef<any>("boxRef")

function fullScreen() {
  if (screenfull.isEnabled) {
    screenfull.toggle(boxRef.value)
  }
}

onMounted(() => {
  monacoCtl = monaco.editor.create(monacoRef.value!, {
    value: "",
    language: "json",
    theme: "vs-dark", //官方自带三种主题vs, hc-black, hc-light, or vs-dark
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
    formatOnPaste: true,
    tabSize: 2,
    folding: true,
    foldingHighlight: true,
    foldingStrategy: "indentation",
    showFoldingControls: "always",
    disableLayerHinting: true,
    emptySelectionClipboard: false,
    selectionClipboard: false,
    codeLens: false,
    colorDecorators: true,
    accessibilitySupport: "off",
    lineNumbers: "on",
    lineNumbersMinChars: 1
  })

  monaco.editor.defineTheme("customTheme", {
    base: "vs", // 基于黑暗主题
    inherit: true, // 继承基础主题
    rules: [],
    colors: {
      // "editor.lineHighlightBackground": "#474747" // 光标行背景色
    }
  })

  // 应用自定义主题
  monaco.editor.setTheme("customTheme")
})

onBeforeUnmount(() => {
  if (monacoCtl) {
    monacoCtl.dispose()
  }
})
</script>

<template>
  <div ref="boxRef" class="monaco-box">
    <div class="monaco-header">
      <el-button link>
        <el-icon :size="20">
          <IconMdiContentCopy />
        </el-icon>
      </el-button>

      <el-button link @click="fullScreen">
        <el-icon :size="20">
          <IconMdiArrowExpandAll />
        </el-icon>
      </el-button>
    </div>
    <div ref="monacoRef" class="monaco-content" />
  </div>
</template>

<style lang="scss" scoped>
.monaco-box {
  --hp-monaco-radius: 4px;
  width: 100%;
  height: 100%;
  border-radius: var(--hp-monaco-radius);
  padding: 4px 8px;
  box-sizing: border-box;
  background-color: #e6e6e7;

  .monaco-header {
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: right;
  }

  .monaco-content {
    width: 100%;
    height: calc(100% - 48px);

    border: 1px dashed #adadad;
    border-radius: var(--hp-monaco-radius);

    :deep(.monaco-editor) {
      border-radius: var(--hp-monaco-radius);
      //--vscode-editor-background: #acb1b0;

      .overflow-guard {
        border-radius: var(--hp-monaco-radius);

        .margin {
          .margin-view-overlays {
            //background-color: #3b3b3b;
          }
        }
      }
    }
  }
}
</style>
