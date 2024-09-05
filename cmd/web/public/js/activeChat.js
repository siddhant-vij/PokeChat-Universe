document.addEventListener("htmx:afterRequest", function () {
  const chatBtns = document.querySelectorAll(".chatBtn");
  chatBtns.forEach((btn) => {
    btn.addEventListener("click", function () {
      const selectedTab = document.querySelector(".active");
      selectedTab.classList.remove("active");
      const chatTab = document.querySelector("#chat-tab");
      chatTab.classList.add("active");
    });
  });
});
