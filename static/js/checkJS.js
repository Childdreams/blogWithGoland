$(".checkboxMain").click(function () {
    if ($(this).is(":checked")){
        $(this).nextAll("input:checkbox").prop("checked",true);
    }else {
        $(this).nextAll("input:checkbox").prop("checked",false);
    }
});

$(".checkboxFork").click(function () {
    var flag = true
    fork = $(this).parent().children(".checkboxFork")
    fork.each(function (i , v) {
        if (!$(v).is(":checked")){
            flag = false
        }
    })
    if (flag){
        $(this).parent().children(".checkboxMain") .prop("checked",true)
    }else {
        $(this).parent().children(".checkboxMain") .prop("checked",false)
    }
})

