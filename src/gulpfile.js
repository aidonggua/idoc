"use strict";
var gulp = require("gulp");
// var rename = require("gulp-rename");
var sass = require("gulp-sass");
// var minifyCss = require("gulp-minify-css");
var browserSync = require("browser-sync").create();
var sourcemaps = require('gulp-sourcemaps');
var webpack = require('webpack');
var gutil = require('gulp-util');

// 公共加载文件生成器
// filename--生成器的文件名，entry Array webpackConfig中的入口key，指定哪些文件使用该加载器加载
// 生成加载器所在的目录为webpackConfig中output的path
var filename = 'loader.js', entry = [];
var CommonChunkPlugin = new webpack.optimize.CommonsChunkPlugin(filename, entry);
var ExtractTextPlugin = require("extract-text-webpack-plugin");


var paths = {
    sass: {
        src: "./scss/**/*.scss",
        output: "./static/css/"
    }
};

gulp.task('sass', function (done) {
    gulp.src(paths.sass.src)
        .pipe(sass())
        .pipe(sourcemaps.init())
        .on('error', sass.logError)
        .pipe(sourcemaps.write())
        .pipe(gulp.dest(paths.sass.output))
        // .pipe(minifyCss({
        //     keepSpecialComments: 0
        // }))
        // .pipe(rename({extname: '.min.css'}))
        // .pipe(gulp.dest(paths.sass.output))
        .on('end', done);
});

gulp.task('browser-sync', function () {
    browserSync.init({
        /* files指定哪些文件发生变更时刷新浏览器 */
        files: ["./template/*.html", "./static/css/*.css", "./static/js/*.js"],
        server: {
            baseDir: "./"
        }
    });
    gulp.watch(paths.sass.src, ['sass']);
    gulp.watch('./**/*.css').on('change', browserSync.reload);
    gulp.watch('./**/*.html').on('change', browserSync.reload);
});

//webpack config
var webpackConfig = {
    plugins: [CommonChunkPlugin, new ExtractTextPlugin('[name].css')],
    // 入口文件
    entry: {
        plugins: ['./js/popup.js', './js/select.js']
    },
    // 编译后的文件 对应entry配置，
    // path:输出文件的目录
    // filename 输出文件的名称，name对应entry的key值
    output: {
        path: 'dist/js/',
        filename: '[name].bundle.js'
    },
    // 指定项目中依赖的jquery来源于外部
    externals: {
        "jquery": "jQuery",
        "bootstrap": ""
    },
    module: {
        loaders: [
            {test: /\.css/, loader: 'style!css'}
        ]
    }
};

gulp.task('webpack', function (done) {
    webpack(webpackConfig, function (err, stats) {
        if (err) {
            throw new gutil.PluginError('webpack', err);
        }
        gutil.log("[webpack]", stats.toString({}));
        done && done();
    })
});