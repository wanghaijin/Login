/**
 * Created by 11293 on 2016.9.6.0006.
 */
//浏览器加载完成后,自动执行此方法
$(function () {
    //获取提交按钮,并绑定单击事件处理程序
    $("#submit").click(function () {
        submit();
    });

    $("#login").click(function () {
        window.location.href = "/";
    });

    //定义提交方法
    function submit() {
        //获取用户名
        var username = $("#username").val();
        //获取密码
        var password = $("#password").val();
        //获取email
        var email = $("#email").val();
        // Email格式
        var emailreg= /^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$/;
        if (username == "") {
            showWarning("请输入用户名!");
            return;
        }
        if (password == "") {
            showWarning("请输入密码!");
            return;
        }
        if(email==""){
            showWarning("请输入邮箱");
            return;
        }
        if(!emailreg.test(email)){
            showWarning("请输入正确的邮箱地址");
            return;
        }

        //提交数据
        $.ajax({
            type: "POST",
            url: "/register",
            data: {"username": username, "password": password, "email": email},
            success: function (data) {
                if (data != "OK") {
                    showWarning(data);
                } else {
                    // showWarning("注册成功！")
                    window.location.href = "/cache";
                }
            },
            error: function () {
                showWarning("注册失败!");
            }
        });

    }


    function showWarning(msg) {
        $("#warning").text(msg);
    }
});


