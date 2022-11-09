interact(".Sign-Img-drag")
  .draggable({
    inertia: true,
    restrict: {
      restriction: "parent",
      endOnly: true,
      elementRect: { top: 0, left: 0, bottom: 1, right: 1 },
    },
    autoScroll: true,
    onmove: handleOnMove,
  })
  .resizable({
    edges: { left: true, right: true, bottom: true, top: true },
    preserveAspectRatio: true,
    restrictEdges: {
      outer: "parent",
      endOnly: true,
    },
    restrictSize: {
      min: { width: 50, height: 25 },
    },
    inertia: true,
  })
  .on("resizemove", handleResizeMove);

function handleOnMove(event) {
  // Event info
  const { target } = event;
  const x = (parseFloat(target.getAttribute("data-x")) || 0) + event.dx;
  const y = (parseFloat(target.getAttribute("data-y")) || 0) + event.dy;

  // Movement
  const transVal = `translate(${x}px, ${y}px)`;
  target.style.transform = target.style.webkitTransform = transVal;
  target.setAttribute("data-x", x);
  target.setAttribute("data-y", y);

  applySignPosition();
}

function handleResizeMove(event) {
  const { target } = event;
  let x = parseFloat(target.getAttribute("data-x")) || 0;
  let y = parseFloat(target.getAttribute("data-y")) || 0;

  // update the element's style
  target.style.width = event.rect.width + "px";
  target.style.height = event.rect.height + "px";

  // translate when resizing from top or left edges
  x += event.deltaRect.left;
  y += event.deltaRect.top;

  const transVal = `translate(${x}px, ${y}px)`;
  target.style.transform = target.style.webkitTransform = transVal;
  target.setAttribute("data-x", x);
  target.setAttribute("data-y", y);

  sign_h = event.rect.height;
  sign_w = event.rect.width;
}

function applySignPosition() {
  // Canvas and signature rect
  // Rect -> element x and y relative to __page__
  const canvEl = document.getElementById("PDFSign");
  const sigEl = document.getElementById("SignImg");
  const canvRect = canvEl.getBoundingClientRect();
  const sigRect = sigEl.getBoundingClientRect();

  // Setting values to be submitted
  // value will be fraction (percentage) of __canvas__
  // Warn: Setting global value
  sign_x = Math.abs(canvRect.left - sigRect.left) / canvRect.width;
  sign_y = Math.abs(canvRect.top - sigRect.top) / canvRect.height;
  sign_h = sigRect.height / canvRect.height;
  sign_w = sigRect.width / canvRect.width;
}

// Put this in global scope to be used in other place
// other place -> clicking 'save' button will call this function
window.applySignPosition = applySignPosition;
