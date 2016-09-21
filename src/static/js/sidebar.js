/**
 * Created by yehao on 16/9/20.
 */

$(function () {
    $('.dir').click(function () {
        $(this).find('span').toggleClass('logo-rotate')
        if ($(this).next('.dir-content').css('display') == "none") {
            $(this).next('.dir-content').show()
        } else {
            $(this).next('.dir-content').hide()
        }
    });
});