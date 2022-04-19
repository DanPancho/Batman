let cargando = document.querySelector('.loading');

setTimeout(()=>{
    cargando.classList.toggle("loading-no");
    bienvenida();
},5000);

const bienvenida = ()=> {
    document.querySelector('.bienvenida-no').classList.toggle("bienvenida");
}