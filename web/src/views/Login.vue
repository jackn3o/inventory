<template>
    <div class="authentication_box">
        <v-layout row
                  align-center
                  justify-center
                  style="height:100vh; align-items: center; overflow: hidden;">
            <v-flex sm12
                    md4
                    d-flex
                    align-center
                    justify-center>

                <v-card :class="contentBoxClasses"
                        style="">
                    <div class="login_form">
                        <div class="alert_box">
                        </div>
                        <v-card-title class="title_box">
                            <!-- <icon name="vanda-logo"></icon>  -->
                            <v-flex sm10
                                    offset-sm1>
                                <div class="display-1 indigo--text">Login</div>
                            </v-flex>
                        </v-card-title>
                        <v-card-text class="mt-3">
                            <v-layout>
                                <v-flex sm10
                                        offset-sm1>
                                    <v-text-field v-model="model.username"
                                                  label="Username"
                                                  prepend-icon="person"
                                                  :error-messages="validations.username"></v-text-field>
                                    <v-text-field v-model="model.password"
                                                  type="password"
                                                  label="Password"
                                                  prepend-icon="lock"
                                                  :error-messages="validations.password"></v-text-field>
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
import validator from '@/mixins/validator'
export default {
    mixins: [validator],
    data() {
        return {
            model: {
                username: null,
                password: null
            }
        }
    },
    computed: {
        contentBoxClasses() {
            return {
                main_content_box: true,
                'elevation-5': this.$vuetify.breakpoint.mdAndUp,
                full_screen: this.$vuetify.breakpoint.smAndDown
            }
        }
    },
    methods: {
        authenticate() {
            let vm = this
            vm.post('/authentication/login', vm.model).then(obj_response => {
                if (!obj_response || !obj_response.data) {
                    return
                }
                this.$store.dispatch('login', obj_response.data)
                this.$router.push({ name: 'inventory.list' })
            })
        }
    }
}
</script>

<style lang="scss">
.authentication_box {
    .main_content_box {
        position: relative;
        border-radius: 4px;
        max-width: 440px;

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

        &.full_screen {
            border-radius: 0;
            max-width: 100%;
            .login_form {
                height: 100vh;
                padding: 16px;
                .title_box {
                    margin-left: -16px;
                }
                .display-1 {
                    margin-left: 16px;
                    font-weight: 700;
                }
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