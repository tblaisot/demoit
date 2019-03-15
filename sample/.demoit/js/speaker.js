window.onload = function(){
    const speakerNotes = document.querySelector(".speaker");
    if (!speakerNotes) return;
    
    const nodes = Array.from(speakerNotes.childNodes);
    nodes.forEach(node => {
        if (node.nodeType === Node.TEXT_NODE) {
            console.log("%c" + node.textContent, window.getComputedStyle(node.parentNode).cssText);
        } else {
            console.log("%c" + node.textContent, window.getComputedStyle(node).cssText);
        }
    })
};