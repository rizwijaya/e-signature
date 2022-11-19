// Toast
$(document).ready(function () {
    $(".toastSuccess").addClass("active");
    $(".progressToast").addClass("active");

    setTimeout(function () { 
        $(".toastSuccess").removeClass("active");
    }, 4000);

    setTimeout(function () { 
        $(".progressToast").removeClass("active");
    }, 4300);

    $(".btn-close").click(function () {
        $(".toastSuccess").removeClass("active");

        setTimeout(function() { 
            $(".progressToast").removeClass("active");
        }, 300);
    })

    $(".toastDanger").addClass("active");
    $(".progressToast").addClass("active");

    setTimeout(function () { 
        $(".toastDanger").removeClass("active");
    }, 4000);

    setTimeout(function () { 
        $(".progressToast").removeClass("active");
    }, 4300);

    $(".btn-close").click(function () {
        $(".toastDanger").removeClass("active");

        setTimeout(function() { 
            $(".progressToast").removeClass("active");
        }, 300);
    })
});

//feather.replace()
// End Script Toast Notifications