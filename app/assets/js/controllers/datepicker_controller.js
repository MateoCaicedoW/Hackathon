import { Controller } from "@hotwired/stimulus";
import flatpickr from "flatpickr";

export default class extends Controller {
  connect() {
    this.datePickerConfig();
  }
  
  dismiss() {
    this.element.classList.add("hidden")
  }

  datePickerConfig() {
    flatpickr(this.element,{
        dateFormat: "m/d/Y",
        defaultDate :"null"
    });

    if (this.element.value == "01/01/0001") {
      this.element._flatpickr.clear();
    }
  }

}
