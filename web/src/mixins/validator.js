const mixin = {
  data() {
    return {
      errorMessage: '',
      responseValidation: null
    }
  },
  computed: {
    showError() {
      return this.errorMessage
    },
    validations() {
      let obj_model = this.model,
        obj_result = {}

      for (let str_key in obj_model) {
        obj_result[str_key] = []
      }

      if (!this.responseValidation) {
        return obj_result
      }

      for (let str_key in obj_result) {
        let str_validation = this.responseValidation[str_key]
        if (str_validation) {
          obj_result[str_key] = [str_validation]
        }
      }

      return obj_result
    }
  },
  methods: {
    handleValidation(obj_error) {
      console.log('validator mixin: start')

      let obj_meta = obj_error.response.data.meta
      if (!obj_meta) {
        console.log('validator mixin: end, no meta')
        return
      }

      let obj_validations = obj_meta.messages.validations
      if (!obj_validations) {
        console.log('validator mixin: end, no validation message')
        return
      }
      this.responseValidation = obj_meta.messages.validations
      console.log('validator mixin: complete')
      return Promise.reject(obj_error)
    },
    resetValidation() {
      this.errorMessage = ''
      this.responseValidation = null
    },
    post(str_url, obj_data = null) {
      this.resetValidation()

      return this.axios({
        method: 'post',
        url: str_url,
        data: obj_data
      }).catch(obj_error => {
        return this.handleValidation(obj_error)
      })
    },
    put(str_url, obj_data = null) {
      this.resetValidation()

      return this.axios({
        method: 'put',
        url: str_url,
        data: obj_data
      }).catch(obj_error => {
        return this.handleValidation(obj_error)
      })
    }
  },
  created() {
    if (!this.model) {
      throw 'validator mixin: missing model'
    }
  }
}

export default mixin
