const wrapper = document.querySelector(".wrapper")
const pagination = document.querySelector(".pagination")
const items = Array.from(document.querySelectorAll(".item"))
let filteredItems = items;
let currPage = 1;

function paginate(totalItems, currentPage = 1, pageSize = 2, maxPages = 3) {

  let totalPages = Math.ceil(totalItems / pageSize);
  if (currentPage < 1) {
    currentPage = 1;
  } else if (currentPage > totalPages) {
    currentPage = totalPages;
  }

  let startPage, endPage;
  if (totalPages <= maxPages) {
    startPage = 1;
    endPage = totalPages;
  } else {
    let maxPagesBeforeCurrentPage = Math.floor(maxPages / 2);
    let maxPagesAfterCurrentPage = Math.ceil(maxPages / 2) - 1;
    if (currentPage <= maxPagesBeforeCurrentPage) {
      startPage = 1;
      endPage = maxPages;
    } else if (currentPage + maxPagesAfterCurrentPage >= totalPages) {
      startPage = totalPages - maxPages + 1;
      endPage = totalPages;
    } else {
      startPage = currentPage - maxPagesBeforeCurrentPage;
      endPage = currentPage + maxPagesAfterCurrentPage;
    }
  }

  let startIndex = (currentPage - 1) * pageSize;
  let endIndex = Math.min(startIndex + pageSize - 1, totalItems - 1);

  let pages = Array.from(Array((endPage + 1) - startPage).keys()).map(i => startPage + i);

  return {
    totalItems: totalItems,
    currentPage: currentPage,
    pageSize: pageSize,
    totalPages: totalPages,
    startPage: startPage,
    endPage: endPage,
    startIndex: startIndex,
    endIndex: endIndex,
    pages: pages
  };
}

function setHTML(items) {
  wrapper.innerHTML = ""
  pagination.innerHTML = ""
  const { currentPage, pageSize, pages } = paginate(items.length, currPage, 2, 3)

  const nav = document.createElement("nav")
  nav.classList.add(...['relative', 'z-0', 'inline-flex', 'rounded-md', 'shadow-sm', '-space-x-px'])

  let paginationHTML = ""

  pages.forEach(page => {
    if (currentPage === page) {
      paginationHTML += `<button class="active z-10 bg-indigo-50 border-indigo-500 text-indigo-600 relative inline-flex items-center px-4 py-2 border text-sm font-medium" page="${page}" ${currentPage === page && 'disabled'}>${page}</button>`
    } else {
      paginationHTML += `<button class="page bg-white border-gray-300 text-gray-500 hover:bg-gray-50 relative inline-flex items-center px-4 py-2 border text-sm font-medium" page="${page}" ${currentPage === page && 'disabled'}>${page}</button>`
    }
  })

  nav.innerHTML = paginationHTML
  pagination.append(nav)

  const start = (currentPage - 1) * pageSize, end = currentPage * pageSize;
  items.slice(start, end).forEach(el => {
    wrapper.append(el)
  })
}

document.body.addEventListener("change", function (e) {
  console.log(e.target);
})
document.addEventListener('click', function (e) {
  const $this = e.target
  console.log($this);
  if ($this.classList.contains("page")) {
    currPage = parseInt($this.getAttribute("page"))
    setHTML(filteredItems)
  }
//   if ($this.classList.contains("next")) {
//     currPage += 1;
//     setHTML(filteredItems)
//   }
//   if ($this.classList.contains("prev")) {
//     currPage -= 1;
//     setHTML(filteredItems)
//   }
});
setHTML(filteredItems)