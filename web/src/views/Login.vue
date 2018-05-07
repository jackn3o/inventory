<template>
    <div class="authentication_box">
        <v-layout row
                  style="height:100vh; align-items: center; overflow: hidden;">
            <v-flex xs12
                    md4
                    offset-md4>
                <v-card height="10"
                        class="small_card">
                </v-card>
                <v-card :class="contentBoxClasses">
                    <div class="login_form">
                        <div class="alert_box">
                        </div>
                        <v-card-title class="title_box">
                            <!-- <icon name="vanda-logo"></icon>  -->
                            <v-flex xs10
                                    offset-xs1>
                                <div class="display-1 indigo--text">Login</div>
                            </v-flex>
                        </v-card-title>
                        <v-card-text class="mt-3">
                            <v-layout>
                                <v-flex xs10
                                        offset-xs1>
                                    <v-text-field label="Username"
                                                  v-model="model.user_name"
                                                  prepend-icon="person"></v-text-field>
                                    <v-text-field type="password"
                                                  label="Password"
                                                  v-model="model.password"
                                                  prepend-icon="lock"></v-text-field>
                                </v-flex>
                            </v-layout>
                        </v-card-text>
                        <v-card-actions class="justify-center">
                            <v-layout row
                                      wrap
                                      style="flex: 1;">
                                <v-flex xs10
                                        md8
                                        offset-xs1
                                        offset-md2>
                                    <v-btn type="button"
                                           class="indigo"
                                           outline
                                           large
                                           block
                                           @click.native="authenticate">SIGN IN</v-btn>
                                </v-flex>
                            </v-layout>
                        </v-card-actions>
                        <div class="alert_box"></div>
                    </div>
                </v-card>
            </v-flex>
        </v-layout>
    </div>

</template>

<script>
export default {
    data() {
        return {
            model: {
                username: null,
                password: null
            }
        }
    },
    computed: {
        r_validation() {
            let obj_model = this.model,
                obj_validation = {}

            for (var str_key in obj_model) {
                obj_validation[str_key] = null
            }

            return r_validation
        },
        contentBoxClasses() {
            let str_viewportSize = this.$store.getters.currentViewportSize

            return {
                main_content_box: true,
                'elevation-5': str_viewportSize != 'xs',
                full_screen: str_viewportSize == 'xs'
            }
        }
    },
    methods: {
        authenticate() {
            let obj_user = {
                username: 'test',
                password: '1111qqqq'
            }
            this.axios.post('localhost:8000/authenticate', obj_user).then(obj_response => {
                console.log(obj_response.data)
            })
        }
    }
}
</script>

<style lang="scss">
.authentication_box {
    .small_card {
        margin: 0 10px;
        border-top-left-radius: 4px;
        border-top-right-radius: 4px;
        border-bottom-left-radius: 0;
        border-bottom-right-radius: 0;
    }

    .main_content_box {
        position: relative;
        border-radius: 4px;
        &.full_screen {
            border-radius: 0;
            .login_form {
                height: 100vh;
            }
            margin-top: -10px;
        }
    }
    .alert_box {
        height: 60px;
        .alert {
            margin: 0;
        }
    }
    .login_form {
        min-height: 480px;
        .title_box {
            border-left: 6px solid #3f51b5;
            .display-1 {
                font-weight: 700;
            }
        }
    }
    .register_form {
        position: absolute;
        top: 0;
        right: 0;
        width: 100%;
        overflow: hidden;
        height: 100%;
        .title_box {
            border-left: 4px solid #fff;
            max-height: 72px;
            .display-1 {
                font-weight: 700;
            }
        }
    }
}
</style>