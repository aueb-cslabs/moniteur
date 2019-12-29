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
        <ErrorPopup ref="error"/>
    </div>
</template>

<script>
    import yaml from 'js-yaml';
    import ErrorPopup from "../ErrorPopup/ErrorPopup";

    const config = require('electron').remote.getGlobal('config');
    const app = require('electron').remote.app;
    const fs = window.require('fs');

    export default {
        name: "Settings",
        components: {ErrorPopup},
        data() {
            return {
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
            checkForm: function (e) {
                if (this.form.api === '') {
                    this.$refs.error.setError(this.$t('message.adminEmptyForm'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.logo_url === '') {
                    this.$refs.error.setError(this.$t('message.adminEmptyForm'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.secondary_logo_url === '') {
                    this.$refs.error.setError(this.$t('message.adminEmptyForm'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.background_color === '') {
                    this.$refs.error.setError(this.$t('message.adminEmptyForm'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.navbar_background_color === '') {
                    this.$refs.error.setError(this.$t('message.adminEmptyForm'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.navbar_color === '') {
                    this.$refs.error.setError(this.$t('message.adminEmptyForm'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                this.save();
                e.preventDefault();
            },

            save: function() {
                fs.writeFileSync(app.getPath('userData')+"/config.yml", yaml.safeDump(this.form), function(err) {
                    if(err) {
                        return err;
                    }
                });
                app.relaunch();
                app.exit(0);
            },
        }
    }
</script>

<style scoped>
@import "../scss/Settings.scss";
</style>