interact('.Sign-Img-drag')
  .draggable({ //drag and drop
    inertia: true,
    restrict: {
      restriction: "parent",
      endOnly: true,
      elementRect: { top: 0, left: 0, bottom: 1, right: 1 }
    },
    autoScroll: true,
    onmove: function (event) {
      var target = event.target,
          x = (parseFloat(target.getAttribute('data-x')) || 0) + event.dx,
          y = (parseFloat(target.getAttribute('data-y')) || 0) + event.dy;

      target.style.webkitTransform = target.style.transform = 'translate(' + x + 'px, ' + y + 'px)';
      target.style.border = '2px dashed #ddd';
      //target.classList.remove('Sign-Img--remove')

      target.setAttribute('data-x', x);
      target.setAttribute('data-y', y);
      //console.log('Coordinate X,Y(' + event.pageX + ', ' + event.pageY + ')')
      sign_x = event.pageX;
      sign_y = event.pageY;
    },
    onend: function (event) {
      var target = event.target;
     // target.classList.add('Sign-Img--remove')
    }
  });
interact('.Sign-Img-drag')
  .resizable({
    // resize from all edges and corners
    edges: { left: true, right: true, bottom: true, top: true },

    //keep aspectratio
    preserveAspectRatio: true,

    // keep the edges inside the parent
    restrictEdges: {
      outer: 'parent',
      endOnly: true,
    },

    // minimum size
    restrictSize: {
      min: { width: 50, height: 25 },
    },

    inertia: true,
  })

  .on('resizemove', function (event) {
    var target = event.target,
        x = (parseFloat(target.getAttribute('data-x')) || 0),
        y = (parseFloat(target.getAttribute('data-y')) || 0);

    // update the element's style
    target.style.width  = event.rect.width + 'px';
    target.style.height = event.rect.height + 'px';

    // console.log('Width: ' + event.rect.width + 'px')
    // console.log('Height: ' + event.rect.height + 'px')
    sign_h = event.rect.height;
    sign_w = event.rect.width;
    // translate when resizing from top or left edges
    x += event.deltaRect.left;
    y += event.deltaRect.top;

    target.style.webkitTransform = target.style.transform =
        'translate(' + x + 'px,' + y + 'px)';
    
    target.setAttribute('data-x', x);
    target.setAttribute('data-y', y);
  });
