// var_1 small entrance animations and focus handling
document.addEventListener('DOMContentLoaded', () => {
  // stagger reveal of .card elements
  const cards = Array.from(document.querySelectorAll('.card'));
  cards.forEach((c, i) => {
    c.style.animationDelay = (0.12 + i * 0.08) + 's';
  });

  // subtle hero pulse
  const hero = document.querySelector('.hero');
  if (hero) {
    setInterval(() => {
      hero.animate([{filter:'brightness(1)'},{filter:'brightness(1.04)'},{filter:'brightness(1)'}], {duration:3200,iterations:1});
    }, 4200);
  }
});
