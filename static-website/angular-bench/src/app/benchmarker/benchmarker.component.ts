import { Component, Input } from '@angular/core';
import { Benchmarker } from '../benchmarker';

@Component({
  selector: 'app-benchmarker',
  standalone: true,
  templateUrl: './benchmarker.component.html',
  styleUrls: ['./benchmarker.component.css']
})
export class BenchmarkerComponent {
  title = 'benchmarker';
  @Input() benchmarker!: Benchmarker;
}
