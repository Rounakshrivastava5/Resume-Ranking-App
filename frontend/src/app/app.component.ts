import { CommonModule, NgFor, NgIf } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterOutlet } from '@angular/router';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, FormsModule, NgFor, NgIf,     CommonModule,
    FormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
    MatProgressSpinnerModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'frontend';
    jobDescription = '';
  resumesRaw = '';
  results: any[] = [];
  loading = false;

  constructor(private http: HttpClient) {
    // You can initialize any 
    // services or perform setup here if needed
  }

  submit() {
    const resumes = this.resumesRaw.split('\n---\n'); // Split resumes
    const body = {
      job_description: this.jobDescription,
      resumes
    };

    this.loading = true;
    this.http.post<any[]>('http://localhost:8080/api/rank', body).subscribe({
      next: res => {
        this.results = res;
        console
        this.loading = false;
      },
      error: err => {
        console.error('Error:', err);
        this.loading = false;
      }
    });
  }
}
