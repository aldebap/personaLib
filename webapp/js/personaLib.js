/* *
     * personaLib init
    */

function personaLibInit() {
}

function personaLibShowAllBooks() {

    //  call book service on the personaLib server
    $.ajax({
        url: "/personaLib/book",
        method: "GET",
        data: {
            csrfmiddlewaretoken: document.getElementsByName("csrfmiddlewaretoken")[0].value
        },
        success: _result => {
            console.log("book list object: \"" + _result + "\"");

            //  crete the list of books
            $("#content").empty();
            $("#content").append("<div id=\"bookList\" class=\"list-group\">");

            books = JSON.parse(_result);
            books.books.forEach(book => {
                $("#bookList").append("<a id=\"" + book.id + "\" href=\"#\" class=\"list-group-item list-group-item-action flex-column align-items-start\">");
                $("#" + book.id).append("<p class=\"mb-1\">" + book.title + "</p>");
                $("#" + book.id).append("<small>" + book.author + "</small>");
            });
        }
    });
}
