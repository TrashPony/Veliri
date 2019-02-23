function anomalyText(text, locateText) {
    Alert(locateText.text, "Неизвестная запись<br>", false, 0, false, "anomalyText");
    let anomalyTextBlock = $('#anomalyText');

    for (let i in locateText.asc) {
        let asc = $('<div></div>');
        asc.addClass('Ask');
        asc.text(locateText.asc[i].text);
        asc.click(function () {
            if (locateText.asc[i].to_page === 0) {
                anomalyTextBlock.remove();
            } else {
                anomalyTextBlock.remove();
                anomalyText(text, text.pages[locateText.asc[i].to_page - 1])
            }
        });
        anomalyTextBlock.append(asc);
    }
}