/**
 * Created by 11293 on 2016.9.6.0006.
 */
$(function () {
    $("#btnReg").click(function () {
        window.location.href="/regedit";
    });

    $("#btnLogin").click(function () {
        login();
    });

    function login() {
        var username=$("#username").val();
        var password=$("#password").val();
        if (username==""){
            showWarning("请输入用户名!");
            return;
        }
        if (password==""){
            showWarning("请输入密码!");
            return;
        }

        $.ajax({
            type:"POST",
            url:"/",
            data:{"username":username,"password":password},
            success:function(data){
                if(data=="OK"){
                    showWarning();
                }else{
                    showWarning();
                }
            },
            error:function () {
                showWarning()

            }
        });
    }

    function  showWarning(msg) {
        $("#warning").text(msg);

    }
})