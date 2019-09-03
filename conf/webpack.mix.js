let mix = require('laravel-mix');

// 禁止系统通知，这个不断通知挺烦人的。
mix.disableNotifications();
// mix.setPublicPath("resources/assets/");

mix.js("/public/static/js/app.js", "static/js/");

mix.sass("/public/static/css/app.scss", "static/css/")