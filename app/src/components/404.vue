<template>
    <div class="container">
        <div class="row">
            <div class="xs-12 md-6 mx-auto">
                <div v-inview:on="countUp">
                    <div class="number">{{count}}</div>
                    <div class="text">Page not found</div>
                    <div class="text">This may not mean anything.</div>
                    <div class="text">I'm probably working on something that has blown up.</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>

    export default {
        name: "notfound",

        data() {
            return {
                count: 0,
                hasRun: false
            }
        },

        methods: {
            formatThousandsNoRounding: function (n, dp) {
                let e = '', s = e + n, l = s.length, b = n < 0 ? 1 : 0,
                    i = s.lastIndexOf(','), j = i == -1 ? l : i,
                    r = e, d = s.substr(j + 1, dp);
                while ((j -= 3) > b) {
                    r = '.' + s.substr(j, 3) + r;
                }
                return s.substr(0, j + 3) + r +
                    (dp ? ',' + d + (d.length < dp ?
                        ('00000').substr(0, dp - d.length) : e) : e);
            },

            countUp($v) {
                $v.enter = () => {
                    if (this.hasRun === false) {
                        this.$el.number.each(function () {
                            let to = (this),
                                countTo = this.count;

                            ({countNum: to.text()}).animate({
                                    countNum: countTo
                                },
                                {
                                    duration: 2000,
                                    easing: 'linear',
                                    step: function () {
                                        to.text(this.formatThousandsNoRounding(Math.floor(this.countNum)));
                                    },
                                    complete: function () {
                                        to.text(this.formatThousandsNoRounding(this.countNum));
                                    }
                                });
                        });
                        this.hasRun = true;
                    }
                }
            }
        }
    }

    /*inView('#countUp').on('enter', function() {

        }
    });*/
</script>

<style>
    @import "../scss/404.scss";
</style>