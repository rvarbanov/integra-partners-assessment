import { ComponentFixture, waitForAsync, TestBed, tick } from '@angular/core/testing';
import { HomeComponent } from './home.component';
import { UserService } from '../services/user.service';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { empty, of, throwError } from 'rxjs';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';

describe('HomeComponent', () => {
  let component: HomeComponent;
  let fixture: ComponentFixture<HomeComponent>;
  let userService: jasmine.SpyObj<UserService>;
  let dialog: jasmine.SpyObj<MatDialog>;
  let snackBar: jasmine.SpyObj<MatSnackBar>;

  const mockUsers = [
    {
      user_id: 1,
      user_name: 'testuser1',
      first_name: 'Test',
      last_name: 'User1',
      email: 'test1@example.com',
      user_status: 'active',
      department: 'IT'
    },
    {
      user_id: 2,
      user_name: 'testuser2',
      first_name: 'Test',
      last_name: 'User2',
      email: 'test2@example.com',
      user_status: 'active',
      department: 'HR'
    }
  ];

  beforeEach(async () => {
    const userServiceSpy = jasmine.createSpyObj('UserService', ['getUsers', 'deleteUser']);
    const dialogSpy = jasmine.createSpyObj('MatDialog', ['open']);
    const snackBarSpy = jasmine.createSpyObj('MatSnackBar', ['open']);

    await TestBed.configureTestingModule({
      imports: [
        HomeComponent,
        NoopAnimationsModule
      ],
      providers: [
        { provide: UserService, useValue: userServiceSpy },
        { provide: MatDialog, useValue: dialogSpy },
        { provide: MatSnackBar, useValue: snackBarSpy }
      ]
    }).compileComponents();

    userService = TestBed.inject(UserService) as jasmine.SpyObj<UserService>;
    dialog = TestBed.inject(MatDialog) as jasmine.SpyObj<MatDialog>;
    snackBar = TestBed.inject(MatSnackBar) as jasmine.SpyObj<MatSnackBar>;
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HomeComponent);
    component = fixture.componentInstance;
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should load users on init', () => {
    userService.getUsers.and.returnValue(of(mockUsers));
    
    fixture.detectChanges();

    expect(userService.getUsers).toHaveBeenCalled();
    expect(component.users).toEqual(mockUsers);
    expect(component.isLoading).toBeFalse();
  });

  it('should handle error when loading users', () => {
    userService.getUsers.and.returnValue(throwError(() => new Error('API Error')));
    component.ngOnInit();
    
    expect(userService.getUsers).toHaveBeenCalled();
    expect(component.users).toEqual([]);
    expect(component.isLoading).toBeFalse();
  });

  it('should delete user', () => {
    const dialogRefSpyObj = jasmine.createSpyObj({ afterClosed: of(true) });
    dialog.open.and.returnValue(dialogRefSpyObj);
    
    userService.deleteUser.and.returnValue(of(undefined));
    
    component.deleteUser(mockUsers[0]);

    expect(dialog.open).toHaveBeenCalled();
    expect(userService.deleteUser).toHaveBeenCalledWith(1);
  });

  // Add more tests as needed
}); 