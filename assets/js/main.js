document.addEventListener('click', (e) => {
    const menu = document.querySelector('[data-user-menu]');
    const checkbox = menu?.querySelector('.app__user-checkbox');

    if (!menu || !checkbox) return;

    if (menu.contains(e.target)) return;

    if (checkbox.checked) {
        checkbox.checked = false;
    }
});
