<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.adminBarSettings")}}</h2>
            <div class="pt-3">
                <form @submit="checkForm">
                    <div class="form-group">
                        <label for="api">API Endpoint</label>
                        <input type="text" class="form-control" id="api" v-model="form.api" placeholder="API Endpoint">
                    </div>
                    <div class="form-group">
                        <label for="logo_url">Logo URL</label>
                        <input type="text" class="form-control" id="logo_url" v-model="form.logo_url" placeholder="Logo URL">
                    </div>
                    <div class="form-group">
                        <label for="sec_logo_url">Secondary Logo URL</label>
                        <input type="text" class="form-control" id="sec_logo_url"
                               v-model="form.secondary_logo_url" placeholder="Secondary Logo URL">
                    </div>
                    <div class="form-group">
                        <label for="background_color">Background color</label>
                        <input type="text" class="form-control" id="background_color" v-model="form.background_color"
                               placeholder="Background color">
                    </div>
                    <div class="form-group">
                        <label for="navbar_background_color">NavBar Background color</label>
                        <input type="text" class="form-control" id="navbar_background_color" v-model="form.navbar_background_color"
                               placeholder="NavBar Background color">
                    </div>
                    <div class="form-group">
                        <label for="navbar_color">NavBar color</label>
                        <input type="text" class="form-control" id="navbar_color" v-model="form.navbar_color"
                               placeholder="NavBar color">
                    </div>
                    <button type="submit" class="btn btn-primary float-right mb-4">{{$t("message.settingsSave")}}</button>
                </form>
            </div>
        </div>
        <div class="error">
            <b-alert
                    :show="dismissCountDown"
                    dismissible
                    variant="danger"
                    @dismissed="dismissCountDown=0"
                    @dismiss-count-down="countDownChanged"
            >
                <p>Error! {{error}}</p>
                <b-progress
                        variant="dark"
                        :max="dismissSecs"
                        :value="dismissCountDown"
                        height="4px"
                ></b-progress>
            </b-alert>
        </div>
    </div>
</template>

<script>
    import yaml from 'js-yaml';

    const config = require('electron').remote.getGlobal('config');
    const fs = window.require('fs');
    //const background = remote.require('../background.js');

    export default {
        name: "Settings",

        data() {
            return {
                dismissSecs: 5,
                dismissCountDown: 0,
                error: null,
                form : {
                    api: config.api,
                    logo_url: config.logo_url,
                    secondary_logo_url: config.secondary_logo_url,
                    background_color: config.background_color,
                    navbar_background_color: config.navbar_background_color,
                    navbar_color: config.navbar_color
                }
            }
        },

        methods: {
            countDownChanged(dismissCountDown) {
                this.dismissCountDown = dismissCountDown
            },

            showAlert() {
                this.dismissCountDown = this.dismissSecs
            },

            checkForm: function (e) {
                if (this.form.api === '') {
                    this.error = this.$t('message.adminEmptyForm');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.logo_url === '') {
                    this.error = this.$t('message.adminEmptyForm');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.secondary_logo_url === '') {
                    this.error = this.$t('message.adminEmptyForm');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.background_color === '') {
                    this.error = this.$t('message.adminEmptyForm');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.navbar_background_color === '') {
                    this.error = this.$t('message.adminEmptyForm');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.navbar_color === '') {
                    this.error = this.$t('message.adminEmptyForm');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                this.error = null;
                this.save();
                e.preventDefault();
            },

            save: function() {
                //saveConfig(this.form);
                fs.writeFileSync("config.yml", yaml.safeDump(this.form), function(err) {
                    if(err) {
                        return err;
                    }
                });
            },
        }
    }
</script>

<style scoped>
@import "../scss/Settings.scss";
</style>