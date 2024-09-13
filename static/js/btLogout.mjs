document.getElementById('btLogout').addEventListener('click', function () {

    window.location.href = '/';

    document.cookie = 'token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';

});