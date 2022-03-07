function if_arrow_give_name_or_return_false(object) {
    if (object.parent.name == "") {
        return false
    } else if (object.parent.name == "z_arrow") {
        return "z"
    } else if (object.parent.name == "y_arrow") {
        return "y"
    } else if (object.parent.name == "x_arrow") {
        return "x"
    } else {
        return false
    }
}

