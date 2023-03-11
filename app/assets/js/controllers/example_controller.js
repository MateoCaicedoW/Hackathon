import {Controller} from  "@hotwired/stimulus"


export default class extends Controller {
    static targets = ["sp"]

    connect(){
        console.log("Hello from Stimulus")
        console.log(this.spTarget);
    }
}