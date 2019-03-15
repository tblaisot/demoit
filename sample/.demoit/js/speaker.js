window.onload = function(){
    let nodes = Array.from(document.querySelector(".speaker").childNodes);
    nodes.forEach(node => {
        if (node.nodeType === Node.TEXT_NODE) {
            console.log("%c" + node.textContent, window.getComputedStyle(node.parentNode).cssText);
        } else {
            console.log("%c" + node.textContent, window.getComputedStyle(node).cssText);
        }
    })
};