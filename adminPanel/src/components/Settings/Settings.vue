<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.adminBarSettings")}}</h2>
            <div class="pt-3">
                <form @submit="save">
                    <SimpleFormGroup label="API Endpoint" v-model="form.api" :model-data="form.api"/>
                    <SimpleFormGroup label="Logo URL" v-model="form.logo_url" :model-data="form.logo_url"/>
                    <SimpleFormGroup label="Secondary Logo URL" v-model="form.secondary_logo_url" :model-data="form.secondary_logo_url"/>
                    <SimpleFormGroup label="Background color" v-model="form.background_color" :model-data="form.background_color" />
                    <SimpleFormGroup label="NavBar Background color" v-model="form.navbar_background_color" :model-data="form.navbar_background_color"/>
                    <SimpleFormGroup label="NavBar color" v-model="form.navbar_color" :model-data="form.navbar_color"/>
                    <button type="submit" class="btn btn-primary float-right mb-4">{{$t("message.settingsSave")}}</button>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
    import yaml from 'js-yaml';
    import SimpleFormGroup from "../SimpleFormGroup/SimpleFormGroup";

    const config = require('electron').remote.getGlobal('config');
    const app = require('electron').remote.app;
    const fs = window.require('fs');

    export default {
        name: "Settings",
        components: {SimpleFormGroup},

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
            save: function(e) {
                fs.writeFileSync(app.getPath('userData')+"/config.yml", yaml.safeDump(this.form), function(err) {
                    if(err) {
                        return err;
                    }
                });
                e.preventDefault();
                app.relaunch();
                app.exit(0);
            },
        }
    }
</script>

<style scoped>
@import "./Settings.scss";
</style>