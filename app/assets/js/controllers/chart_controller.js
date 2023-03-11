import { Controller } from "@hotwired/stimulus";
import Chart from "chart.js/auto";
import { getRelativePosition } from 'chart.js/helpers';

export default class extends Controller {
    static values = {
        result: String,
    }

  connect() {
    super.connect();
    this.chart();
  }
     
}
