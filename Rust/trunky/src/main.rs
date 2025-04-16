use web_sys::{window, Text, HtmlElement};

fn main() {
    console_error_panic_hook::set_once();

    let document = window()
        .and_then(|win| win.document())
        .expect("Could not access the document");

    let body = document.body().expect("Could not access document.body");

    // First <p>
    let p1 = document.create_element("p").expect("Failed to create <p>");
    let text_node1: Text = document.create_text_node("Hello, Old Chap!");
    p1.append_child(&text_node1).expect("Failed to append text to <p1>");
    body.append_child(&p1).expect("Failed to append <p1> to body");

    // Second <p>
    let p2 = document.create_element("p").expect("Failed to create <p>");
    let text_node2: Text = document.create_text_node("This is it!");
    p2.append_child(&text_node2).expect("Failed to append text to <p2>");
    body.append_child(&p2).expect("Failed to append <p2> to body");
}
