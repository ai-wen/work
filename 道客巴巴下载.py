function downloadPages(from, to) {
    for (i = from; i <= to; i++) {
        const pageCanvas = document.getElementById('page_' + i);
        if (pageCanvas === null) { break; }
        const pageNo = parseInt(String(i));
        setTimeout(() => {
            console.log("==pageNo==>>", pageNo);
            ((num) => {
                console.log("开始打印第" + num + "页");
                pageCanvas.toBlob(
                    blob => {
                        const anchor = document.createElement('a');
                        anchor.download = 'page_' + num + '.png';
                        anchor.href = URL.createObjectURL(blob);
                        anchor.click();
                        URL.revokeObjectURL(anchor.href);
                    }
                );
            })(pageNo);
        }, 500 * pageNo);
    }
}

//浏览器命令行执行，下载图片
downloadPages(1,59);