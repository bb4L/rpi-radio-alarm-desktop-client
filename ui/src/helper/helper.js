
function toast(ref, title, content, toaster = "b-toaster-bottom-right", variant = "info", append = true, noAutoHide = false) {
    ref.$bvToast.toast(content, {
        title: title,
        toaster: toaster,
        solid: true,
        appendToast: append,
        variant: variant,
        noAutoHide: noAutoHide,
        autoHideDelay: 1000,
    });
}

function infoToast(ref, title, content, toaster = "b-toaster-bottom-right", append = true, noAutoHide = false) {
    toast(ref, title, content, toaster, "info", append, noAutoHide)
}

function successToast(ref, title, content, toaster = "b-toaster-bottom-right", append = true, noAutoHide = false) {
    toast(ref, title, content, toaster, "success", append, noAutoHide)
}

function errorToast(ref, title, content, toaster = "b-toaster-bottom-right", append = true, noAutoHide = true) {
    toast(ref, title, content, toaster, "danger", append, noAutoHide)
}

function delay(t) {
    return new Promise((resolve) => setTimeout(resolve, t));
}


export default {
    toast,
    infoToast,
    successToast,
    errorToast,
    delay
}