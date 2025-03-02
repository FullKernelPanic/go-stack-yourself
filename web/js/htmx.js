import htmx from "htmx.org";

export default function initHTMX() {
    window.htmx = htmx;

    window.htmx.config.responseHandling = [
        {code: "204", swap: false},   // 204 - No Content by default does nothing, but is not an error
        {code: "[23]..", swap: true}, // 200 & 300 responses are non-errors and are swapped
        {code: "[4]..", swap: true, error: true}, // 400 & 500 responses are not swapped and are errors
    ];

    htmx.config.logger = function(elt, event, data) {
        console.log(`HTMX event: ${event}`, {element: elt, data: data});
    };
}