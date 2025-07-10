-- CreateTable
CREATE TABLE "role" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "description" TEXT DEFAULT '',
    "is_active" BOOLEAN DEFAULT true,
    "permissions" JSONB DEFAULT [],
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL
);

-- CreateTable
CREATE TABLE "permission" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL
);

-- CreateTable
CREATE TABLE "user" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "role_id" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "image" TEXT,
    "contacts" JSONB DEFAULT [],
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    "is_active" BOOLEAN NOT NULL DEFAULT true,
    CONSTRAINT "user_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "role" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "customer" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "title" TEXT,
    "first_name" TEXT,
    "middle_name" TEXT,
    "last_name" TEXT,
    "date_of_birth" TEXT,
    "contacts" JSONB DEFAULT [],
    "is_lead" BOOLEAN NOT NULL DEFAULT false,
    "description" TEXT DEFAULT '',
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL
);

-- CreateTable
CREATE TABLE "provider" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "description" TEXT DEFAULT '',
    "contacts" JSONB DEFAULT [],
    "is_active" BOOLEAN NOT NULL DEFAULT true,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL
);

-- CreateTable
CREATE TABLE "cash_register" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "owner_id" TEXT NOT NULL,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    CONSTRAINT "cash_register_owner_id_fkey" FOREIGN KEY ("owner_id") REFERENCES "user" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "ledger_sessions" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "ledger_id" TEXT NOT NULL,
    "opened_by" TEXT NOT NULL,
    "closed_by" TEXT,
    "opening_amount" JSONB NOT NULL,
    "closing_amount" JSONB NOT NULL,
    "status" TEXT NOT NULL,
    "notes" TEXT,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    CONSTRAINT "ledger_sessions_ledger_id_fkey" FOREIGN KEY ("ledger_id") REFERENCES "cash_register" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "ledger_sessions_opened_by_fkey" FOREIGN KEY ("opened_by") REFERENCES "user" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "ledger_sessions_closed_by_fkey" FOREIGN KEY ("closed_by") REFERENCES "user" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "movement" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "session_id" TEXT NOT NULL,
    "registered_by" TEXT NOT NULL,
    "amount" DECIMAL NOT NULL,
    "type" TEXT NOT NULL,
    "description" TEXT,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    CONSTRAINT "movement_session_id_fkey" FOREIGN KEY ("session_id") REFERENCES "ledger_sessions" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "movement_registered_by_fkey" FOREIGN KEY ("registered_by") REFERENCES "user" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "destination" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "name" JSONB NOT NULL,
    "description" JSONB NOT NULL,
    "parent_id" TEXT,
    "location" JSONB,
    "is_deleted" BOOLEAN NOT NULL DEFAULT false,
    "is_active" BOOLEAN NOT NULL DEFAULT true,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    CONSTRAINT "destination_parent_id_fkey" FOREIGN KEY ("parent_id") REFERENCES "destination" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "tag" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "title" JSONB NOT NULL,
    "description" JSONB NOT NULL,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL
);

-- CreateTable
CREATE TABLE "resource_type" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "type" TEXT NOT NULL
);

-- CreateTable
CREATE TABLE "resource" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "destination_id" TEXT,
    "location" JSONB,
    "name" JSONB NOT NULL,
    "description" JSONB NOT NULL,
    "images" JSONB NOT NULL DEFAULT [],
    "is_active" BOOLEAN NOT NULL DEFAULT true,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL
);

-- CreateTable
CREATE TABLE "resource_resource_type" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "resource_id" TEXT NOT NULL,
    "type_id" TEXT NOT NULL,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    CONSTRAINT "resource_resource_type_resource_id_fkey" FOREIGN KEY ("resource_id") REFERENCES "resource" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "resource_resource_type_type_id_fkey" FOREIGN KEY ("type_id") REFERENCES "resource_type" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "resource_provider" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "resource_id" TEXT NOT NULL,
    "provider_id" TEXT,
    "user_id" TEXT,
    "ref_prices" JSONB DEFAULT [],
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    CONSTRAINT "resource_provider_resource_id_fkey" FOREIGN KEY ("resource_id") REFERENCES "resource" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "resource_provider_provider_id_fkey" FOREIGN KEY ("provider_id") REFERENCES "provider" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT "resource_provider_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "tour" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "code" TEXT NOT NULL,
    "name" JSONB NOT NULL,
    "slug" JSONB NOT NULL,
    "discount_percent" DECIMAL NOT NULL DEFAULT 1,
    "days" INTEGER NOT NULL DEFAULT 0,
    "group_size" TEXT NOT NULL DEFAULT '',
    "transport" JSONB NOT NULL,
    "accommodation" JSONB NOT NULL,
    "team" JSONB NOT NULL,
    "short_description" JSONB NOT NULL,
    "long_description" JSONB NOT NULL,
    "images" JSONB NOT NULL DEFAULT [],
    "is_public" BOOLEAN NOT NULL DEFAULT false,
    "is_deleted" BOOLEAN NOT NULL DEFAULT false,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL
);

-- CreateTable
CREATE TABLE "tour_tag" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "tour_id" TEXT NOT NULL,
    "tag_id" TEXT NOT NULL,
    CONSTRAINT "tour_tag_tour_id_fkey" FOREIGN KEY ("tour_id") REFERENCES "tour" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "tour_tag_tag_id_fkey" FOREIGN KEY ("tag_id") REFERENCES "tag" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "tour_variant" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "tour_id" TEXT NOT NULL,
    "min_persons" INTEGER NOT NULL,
    "max_persons" INTEGER NOT NULL,
    "totalPrice" DECIMAL NOT NULL DEFAULT 0,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL
);

-- CreateTable
CREATE TABLE "section_description" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "tour_variant_id" TEXT NOT NULL,
    "tour_id" TEXT NOT NULL,
    "images" JSONB NOT NULL DEFAULT [],
    "description" JSONB NOT NULL,
    "startDate" INTEGER NOT NULL DEFAULT 0,
    "endDate" INTEGER NOT NULL DEFAULT 0,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    CONSTRAINT "section_description_tour_id_fkey" FOREIGN KEY ("tour_id") REFERENCES "tour" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "section_description_tour_variant_id_fkey" FOREIGN KEY ("tour_variant_id") REFERENCES "tour_variant" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "tour_destination" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "destination_id" TEXT NOT NULL,
    "tour_id" TEXT NOT NULL,
    "tour_variant_id" TEXT NOT NULL,
    "description" JSONB NOT NULL,
    "startDate" INTEGER NOT NULL DEFAULT 0,
    "endDate" INTEGER NOT NULL DEFAULT 0,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    CONSTRAINT "tour_destination_destination_id_fkey" FOREIGN KEY ("destination_id") REFERENCES "destination" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "tour_destination_tour_id_fkey" FOREIGN KEY ("tour_id") REFERENCES "tour" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "tour_destination_tour_variant_id_fkey" FOREIGN KEY ("tour_variant_id") REFERENCES "tour_variant" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "tour_resource_provider" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "tour_id" TEXT NOT NULL,
    "tour_variant_id" TEXT NOT NULL,
    "resource_provider_id" TEXT NOT NULL,
    "startDate" INTEGER NOT NULL DEFAULT 0,
    "endDate" INTEGER NOT NULL DEFAULT 0,
    "original_cost" DECIMAL,
    "currency" TEXT,
    "profit_percent" DECIMAL,
    "dollar_change_rage" DECIMAL,
    "quantity" INTEGER NOT NULL DEFAULT 1,
    "is_visible" BOOLEAN NOT NULL DEFAULT false,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    CONSTRAINT "tour_resource_provider_tour_id_fkey" FOREIGN KEY ("tour_id") REFERENCES "tour" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "tour_resource_provider_resource_provider_id_fkey" FOREIGN KEY ("resource_provider_id") REFERENCES "resource_provider" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "tour_resource_provider_tour_variant_id_fkey" FOREIGN KEY ("tour_variant_id") REFERENCES "tour_variant" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "trip" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "tour_id" TEXT NOT NULL,
    "registered_by" TEXT NOT NULL,
    "customer_id" TEXT NOT NULL,
    "state" TEXT NOT NULL DEFAULT 'programmed',
    "start_date" INTEGER NOT NULL,
    "end_date" INTEGER NOT NULL,
    "available_slots" INTEGER NOT NULL,
    "num_guests" INTEGER NOT NULL,
    "price_pp" DECIMAL NOT NULL,
    "total" DECIMAL NOT NULL,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    CONSTRAINT "trip_tour_id_fkey" FOREIGN KEY ("tour_id") REFERENCES "tour" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "trip_registered_by_fkey" FOREIGN KEY ("registered_by") REFERENCES "user" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "trip_customer_id_fkey" FOREIGN KEY ("customer_id") REFERENCES "customer" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "payment" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "trip_id" TEXT NOT NULL,
    "customer_id" TEXT NOT NULL,
    "registered_by" TEXT NOT NULL,
    "movement_id" TEXT,
    "amount" DECIMAL NOT NULL,
    "payment_method" TEXT NOT NULL,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    CONSTRAINT "payment_trip_id_fkey" FOREIGN KEY ("trip_id") REFERENCES "trip" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "payment_customer_id_fkey" FOREIGN KEY ("customer_id") REFERENCES "customer" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "payment_registered_by_fkey" FOREIGN KEY ("registered_by") REFERENCES "user" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "payment_movement_id_fkey" FOREIGN KEY ("movement_id") REFERENCES "movement" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "trip_expense" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "created_by_id" TEXT NOT NULL,
    "trip_id" TEXT NOT NULL,
    "movement_id" TEXT,
    "amount" DECIMAL NOT NULL,
    "description" TEXT NOT NULL,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    CONSTRAINT "trip_expense_trip_id_fkey" FOREIGN KEY ("trip_id") REFERENCES "trip" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "trip_expense_created_by_id_fkey" FOREIGN KEY ("created_by_id") REFERENCES "user" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "trip_expense_movement_id_fkey" FOREIGN KEY ("movement_id") REFERENCES "movement" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "trip_destination" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "trip_id" TEXT NOT NULL,
    "destination_id" TEXT NOT NULL,
    "price_pp" DECIMAL NOT NULL,
    "start_date" INTEGER NOT NULL,
    "end_date" INTEGER NOT NULL,
    CONSTRAINT "trip_destination_trip_id_fkey" FOREIGN KEY ("trip_id") REFERENCES "trip" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "trip_destination_destination_id_fkey" FOREIGN KEY ("destination_id") REFERENCES "destination" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "trip_resource" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "trip_id" TEXT NOT NULL,
    "resource_id" TEXT NOT NULL,
    "movement_id" TEXT,
    "total" DECIMAL NOT NULL,
    "start_date" INTEGER NOT NULL,
    "end_date" INTEGER NOT NULL,
    CONSTRAINT "trip_resource_trip_id_fkey" FOREIGN KEY ("trip_id") REFERENCES "trip" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "trip_resource_resource_id_fkey" FOREIGN KEY ("resource_id") REFERENCES "resource" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "trip_resource_movement_id_fkey" FOREIGN KEY ("movement_id") REFERENCES "movement" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);

-- CreateIndex
CREATE UNIQUE INDEX "role_name_key" ON "role"("name");

-- CreateIndex
CREATE UNIQUE INDEX "user_email_key" ON "user"("email");

-- CreateIndex
CREATE UNIQUE INDEX "tour_code_key" ON "tour"("code");
