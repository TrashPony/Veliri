let action = {event: ''};
action.zoom = Zoom;

async function Zoom() {
    function sleep(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

    while(this.event !== ''){
        if (this.event === 'zoomUp') {
            SizeGameMap(+0.001)
        } else if (this.event === 'zoomDown') {
            SizeGameMap(-0.001)
        }
        await sleep(10);
    }
}

function mouseZoomPress(needEvent) {
    action.event = needEvent;
    action.zoom();
}