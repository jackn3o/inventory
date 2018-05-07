const helper = {
  methods: {
    isStringContain(value, key) {
      if (!value || !key) {
        return false
      }

      let str_value = String(value).toLowerCase(),
        str_key = String(key).toLowerCase()

      if (str_value.search(str_key) >= 0) {
        return true
      }

      return false
    },
  }
}

export default helper
