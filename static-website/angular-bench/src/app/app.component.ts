import { Component } from '@angular/core';
import { BenchmarkerComponent } from './benchmarker/benchmarker.component';
import { Benchmarker } from './benchmarker';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    BenchmarkerComponent,
  ],
  template: `
  <main>
  <section class="content">
    <app-benchmarker [benchmarker]="benchmarker"></app-benchmarker>
  </section>
</main>
  `,
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  currentTime = new Date();
  //@ts-ignore
  timer = this.currentTime - window.performance.timing.requestStart;
  title = 'angular-bench';

  benchmarker: Benchmarker = {
    timer: this.timer
  }
}
