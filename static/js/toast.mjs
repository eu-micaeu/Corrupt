function toastGreen(message) {

    const toast = document.createElement('div');

    toast.textContent = message;

    toast.style.position = 'fixed';

    toast.style.bottom = '20px';

    toast.style.right = '20px';

    toast.style.padding = '10px 20px';

    toast.style.backgroundColor = '#4CAF50';

    toast.style.color = '#fff';

    toast.style.borderRadius = '5px';

    toast.style.boxShadow = '0 4px 8px rgba(0, 0, 0, 0.1)';

    toast.style.fontSize = '16px';

    toast.style.opacity = '0';

    toast.style.transition = 'opacity 0.3s ease, transform 0.3s ease';

    toast.style.transform = 'translateY(20px)';

    document.body.appendChild(toast);

    setTimeout(() => {

        toast.style.opacity = '1';

        toast.style.transform = 'translateY(0)';

    }, 100);

    setTimeout(() => {

        toast.style.opacity = '0';

        toast.style.transform = 'translateY(20px)';

        setTimeout(() => {

            toast.remove();

        }, 300);

    }, 3000);

}

function toastRed(message) {

    const toast = document.createElement('div');

    toast.textContent = message;

    toast.style.position = 'fixed';

    toast.style.bottom = '20px';

    toast.style.right = '20px';

    toast.style.padding = '10px 20px';

    toast.style.backgroundColor = '#f44336';

    toast.style.color = '#fff';

    toast.style.borderRadius = '5px';

    toast.style.boxShadow = '0 4px 8px rgba(0, 0, 0, 0.1)';

    toast.style.fontSize = '16px';

    toast.style.opacity = '0';

    toast.style.transition = 'opacity 0.3s ease, transform 0.3s ease';

    toast.style.transform = 'translateY(20px)';

    document.body.appendChild(toast);

    setTimeout(() => {

        toast.style.opacity = '1';

        toast.style.transform = 'translateY(0)';

    }, 100);

    setTimeout(() => {

        toast.style.opacity = '0';

        toast.style.transform = 'translateY(20px)';

        setTimeout(() => {

            toast.remove();

        }, 300);

    }, 3000);

}

export { toastGreen, toastRed };