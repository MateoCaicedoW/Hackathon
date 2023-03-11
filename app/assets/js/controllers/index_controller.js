import { Controller } from  "@hotwired/stimulus"


export default class extends Controller {

    static targets = [
        "navigation",
        "icon"
    ]


    changeIcon(e){
        let event = e.target;  
        event.innerHTML === "close" ? event.innerHTML = "menu" : event.innerHTML = "close";
        this.navigationTarget.classList.toggle("hidden");
    }


    closeNavigation(e){

        if (e.target.classList.contains("navigation")) {
            return
        }
        this.navigationTarget.classList.add("hidden");
        this.iconTarget.innerHTML = "menu"; 
        
    }
    
    
}