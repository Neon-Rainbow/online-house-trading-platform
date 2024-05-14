$(document).ready(function() {
    // 当选择器（radio buttons）的状态改变时，执行函数
    $("input[name='gender']").change(function() {
        // 获取选择器（radio buttons）中被选中的值
        var gender = $("input[name='gender']:checked").val();
        
        // 输出选中的值
        console.log("性别:", gender);
    });
});


$(document).ready(function() {
    $("#personal-messages-save-btn").click(function() {
        var  personalemail= $("#personal-email-input").val();
        var  Verificationcode= $("#Verification-code-input").val();
        var personalpasswordinput = $("#personal-password-input").val();
        var  confirminput= $("#personal-password-confirm-input").val();
        var  Nickname= $("#personal-Nickname-input").val();
        var  phonenumber= $("#personal-phonenumber-input").val();
        
        $.ajax({
            url: "/api/save-personal-info",
            type: "POST",
            dataType: "json",
            data: {
                personalemail:personalemail,
                Verificationcode:Verificationcode,
                
            },
            success: function (resp) {
                console.log("个人信息保存成功:", resp);
                alert("个人信息保存成功！");
            },
            error: function (xhr, status, error) {
                console.error("保存个人信息失败:", error);
                alert("保存个人信息失败，请稍后重试！");
            }
        });
        
        // 发送保存请求
        
    });
});