export default defineStore("page", () => {
  const screenWidth = ref(window.innerWidth)
  const screenHeight = ref(window.innerHeight)
  const menuIsCollapse = ref(false)

  const refreshPage = ref({
    needRefresh: false,
    type: ""
  })

  const isBigScreen = computed(() => {
    return screenWidth.value > 1200
  })

  const isSmallScreen = computed(() => {
    return screenWidth.value < 800
  })

  function setScreen() {
    screenWidth.value = window.innerWidth
    screenHeight.value = window.innerHeight
  }

  function changeCollapse() {
    menuIsCollapse.value = !menuIsCollapse.value
  }

  return {
    screenWidth,
    screenHeight,
    isBigScreen,
    isSmallScreen,
    refreshPage,
    menuIsCollapse,
    setScreen,
    changeCollapse
  }
})
