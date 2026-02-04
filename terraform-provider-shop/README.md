# 🪟 Windows CMD – Execute `shop_customer` Resource

This guide walks through **building, installing, and executing** the `shop_customer` Terraform resource on **Windows (CMD)**.

---

# IN SHORT :- OPEN CMD
```
cd D:\Work_dsi\MyPersonalProjects\GoLang\terraform-provider-shop
go mod tidy
go build -o terraform-provider-shop.exe

mkdir %APPDATA%\terraform.d\plugins\registry.terraform.io\yashgaykar\shop\0.1.0\windows_amd64
copy terraform-provider-shop.exe %APPDATA%\terraform.d\plugins\registry.terraform.io\yashgaykar\shop\0.1.0\windows_amd64\

cd examples
terraform init
terraform apply
```


## 1️⃣ Build the Provider (CMD)

Open **Command Prompt** and run:

```cmd
cd D:\Work_dsi\MyPersonalProjects\GoLang\terraform-provider-shop
go mod tidy
go build -o terraform-provider-shop.exe
```

### Expected

- ✅ No errors
- 📦 File created: `terraform-provider-shop.exe`

---

## 2️⃣ Install Provider Locally (Windows CMD)

### Terraform plugin directory on Windows

```cmd
mkdir %APPDATA%\terraform.d\plugins\registry.terraform.io\yashgaykar\shop\0.1.0\windows_amd64
```

e.g  C:\Users\YGR2\AppData\Roaming\terraform.d\plugins\registry.terraform.io\yashgaykar\shop\0.1.0\windows_amd64 

### Copy the provider binary

```cmd
copy terraform-provider-shop.exe %APPDATA%\terraform.d\plugins\registry.terraform.io\yashgaykar\shop\0.1.0\windows_amd64\
```

---

## 3️⃣ Verify Shop Application Is Running

```cmd
curl http://localhost:8080/customers
```

### Expected

- Any valid response = ✅ OK

---

## 4️⃣ Run Terraform Example

```cmd
cd examples
terraform init
terraform apply
```

When prompted, type:

```text
yes
```

### Expected

- ✅ Customer created via API
- ✅ Terraform state updated

---

## 5️⃣ Validate Lifecycle (CMD)

### Re-plan (no changes expected)

```cmd
terraform plan
```

### Update test

1. Change `address` in `customer.tf`
2. Apply again:

```cmd
terraform apply
```

### Destroy

```cmd
terraform destroy
```

---

## 🛑 If Terraform Can’t Find the Provider (Fix)

Run once:

```cmd
terraform init -upgrade
```

If still stuck, delete and re-init:

```cmd
rmdir /s /q .terraform
terraform init
```

---

## ✅ Stop Point (As Agreed)

Stop once all of the following are true:

- `terraform apply` works
- `terraform plan` is clean
- `terraform destroy` works

👉 **Stop for today.** Everything else can wait.

---

### 🆘 If You Hit Any Error

Paste **only the error output**, nothing else — I’ll decode