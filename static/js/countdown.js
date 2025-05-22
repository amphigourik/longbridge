document.addEventListener('DOMContentLoaded', () => {
    const targetDate = new Date("2026-04-25T15:00:00"); // Date du mariage
    const countdownEl = document.getElementById("countdown");
  
    function updateCountdown() {
        const now = new Date();
        const diff = targetDate - now;
    
        if (diff <= 0) {
            countdownEl.textContent = "C'est aujourd'hui ❤️";
            return;
        }
    
        const days = Math.floor(diff / (1000 * 60 * 60 * 24));
        const hours = Math.floor((diff / (1000 * 60 * 60)) % 24);
        const minutes = Math.floor((diff / (1000 * 60)) % 60);
        const seconds = Math.floor((diff / 1000) % 60);
    
        countdownEl.textContent = `${days}j ${hours}h ${minutes}m ${seconds}s`;
    }
  
    updateCountdown();
    setInterval(updateCountdown, 1000);
});
  